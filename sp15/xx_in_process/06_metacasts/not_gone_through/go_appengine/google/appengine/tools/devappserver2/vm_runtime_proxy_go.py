#!/usr/bin/env python
#
# Copyright 2007 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
"""Manages a Go VM Runtime process running inside of a docker container.
"""

import logging
import os
import shutil
import tempfile

import google

from google.appengine.tools.devappserver2 import go_application
from google.appengine.tools.devappserver2 import instance
from google.appengine.tools.devappserver2 import vm_runtime_proxy

REQUEST_ID_HEADER_NAME = 'X-Appengine-Api-Ticket'
DEBUG_PORT = 5858
VM_SERVICE_PORT = 8181
# TODO: Remove this when classic Go SDK is gone.
DEFAULT_DOCKER_FILE = """FROM google/appengine-go

ADD . /app
RUN /bin/bash /app/_ah/build.sh
"""

# Where to look for go-app-builder, which is needed for copying
# into the Docker image for building the Go App Engine app.
# There is no need to add '.exe' here because it is always a Linux executable.
_GO_APP_BUILDER = os.path.join(
    go_application.GOROOT, 'pkg', 'tool', 'docker-gab')


class GoVMRuntimeProxy(instance.RuntimeProxy):
  """Manages a Go VM Runtime process running inside of a docker container.

  The Go VM Runtime forwards all requests to the Go application instance.
  """

  def __init__(self, docker_client, runtime_config_getter,
               module_configuration):
    """Initializer for VMRuntimeProxy.

    Args:
      docker_client: docker.Client object to communicate with Docker daemon.
      runtime_config_getter: A function that can be called without arguments
          and returns the runtime_config_pb2.Config containing the configuration
          for the runtime.
      module_configuration: An application_configuration.ModuleConfiguration
          instance respresenting the configuration of the module that owns the
          runtime.
    """
    super(GoVMRuntimeProxy, self).__init__()
    self._runtime_config_getter = runtime_config_getter
    self._module_configuration = module_configuration
    port_bindings = {
        DEBUG_PORT: None,
        VM_SERVICE_PORT: None,
    }
    self._vm_runtime_proxy = vm_runtime_proxy.VMRuntimeProxy(
        docker_client=docker_client,
        runtime_config_getter=runtime_config_getter,
        module_configuration=module_configuration,
        port_bindings=port_bindings)

  def handle(self, environ, start_response, url_map, match, request_id,
             request_type):
    """Handle request to Go runtime.

    Serves this request by forwarding to the Go application instance via
    HttpProxy.

    Args:
      environ: An environ dict for the request as defined in PEP-333.
      start_response: A function with semantics defined in PEP-333.
      url_map: An appinfo.URLMap instance containing the configuration for the
          handler matching this request.
      match: A re.MatchObject containing the result of the matched URL pattern.
      request_id: A unique string id associated with the request.
      request_type: The type of the request. See instance.*_REQUEST module
          constants.

    Yields:
      A sequence of strings containing the body of the HTTP response.
    """

    it = self._vm_runtime_proxy.handle(environ, start_response, url_map,
                                       match, request_id, request_type)
    for data in it:
      yield data

  def start(self):
    logging.info('Starting Go VM Deployment process')

    try:
      application_dir = os.path.abspath(
          self._module_configuration.application_root)

      with TempDir('go_deployment_dir') as dst_deployment_dir:
        build_go_docker_image_source(
            application_dir, dst_deployment_dir, _GO_APP_BUILDER,
            self._module_configuration.nobuild_files,
            self._module_configuration.skip_files)

        self._vm_runtime_proxy.start(
            dockerfile_dir=dst_deployment_dir,
            request_id_header_name=REQUEST_ID_HEADER_NAME)

      logging.info(
          'GoVM vmservice for module "%(module)s" available at '
          'http://%(host)s:%(port)s/',
          {
              'module': self._module_configuration.module_name,
              'host': self._vm_runtime_proxy.ContainerHost(),
              'port': self._vm_runtime_proxy.PortBinding(VM_SERVICE_PORT),
          })

    except Exception as e:
      logging.info('Go VM Deployment process failed: %s', str(e))
      raise

  def quit(self):
    self._vm_runtime_proxy.quit()


def _write_dockerfile(dst_dir):
  """Writes Dockerfile to named directory if one does not exist.

  Args:
    dst_dir: string name of destination directory.
  """
  dst_dockerfile = os.path.join(dst_dir, 'Dockerfile')
  if not os.path.exists(dst_dockerfile):
    with open(dst_dockerfile, 'w') as fd:
      fd.write(DEFAULT_DOCKER_FILE)


class TempDir(object):
  """Creates a temporary directory."""

  def __init__(self, prefix=''):
    self._temp_dir = None
    self._prefix = prefix

  def __enter__(self):
    self._temp_dir = tempfile.mkdtemp(self._prefix)
    return self._temp_dir

  def __exit__(self, *_):
    shutil.rmtree(self._temp_dir, ignore_errors=True)


def _copytree(src, dst, skip_files, symlinks=False):
  """Copies src tree to dst (except those matching skip_files).

  Args:
    src: string name of source directory to copy from.
    dst: string name of destination directory to copy to.
    skip_files: RegexStr of files to skip from appinfo.py.
    symlinks: optional bool determines if symbolic links are followed.
  """
  # Ignore files that match the skip_files RegexStr.
  # TODO: skip_files expects the full path relative to the app root, so
  # this may need fixing.
  def ignored_files(unused_dir, filenames):
    return [filename for filename in filenames if skip_files.match(filename)]

  for item in os.listdir(src):
    s = os.path.join(src, item)
    if skip_files.match(item):
      logging.info('skipping file %s', s)
      continue
    d = os.path.join(dst, item)
    if os.path.isdir(s):
      shutil.copytree(s, d, symlinks, ignore=ignored_files)
    else:
      shutil.copy2(s, d)


def build_go_docker_image_source(
    application_dir, dst_deployment_dir, go_app_builder,
    nobuild_files, skip_files):
  """Builds the Docker image source in preparation for building.

  Steps:
    copy the application to dst_deployment_dir (follow symlinks)
    copy used parts of $GOPATH to dst_deployment_dir
    copy or create a Dockerfile in dst_deployment_dir

  Args:
    application_dir: string pathname of application directory.
    dst_deployment_dir: string pathname of temporary deployment directory.
    go_app_builder: string pathname of docker-gab executable.
    nobuild_files: regexp identifying which files to not build.
    skip_files: regexp identifying which files to omit from app.
  """
  try:
    _copytree(application_dir, dst_deployment_dir, skip_files)
  except shutil.Error as e:
    logging.error('Error copying tree: %s', e)
    for src, unused_dst, unused_error in e.args[0]:
      if os.path.islink(src):
        linkto = os.readlink(src)
        if not os.path.exists(linkto):
          logging.error('Dangling symlink in Go project. '
                        'Path %s links to %s', src, os.readlink(src))
    raise
  except OSError as e:
    logging.error('Failed to copy dir: %s', e.strerror)
    raise

  extras = go_application.get_app_extras_for_vm(
      application_dir, nobuild_files, skip_files)
  for dest, src in extras:
    try:
      dest = os.path.join(dst_deployment_dir, dest)
      dirname = os.path.dirname(dest)
      if not os.path.exists(dirname):
        os.makedirs(dirname)
      shutil.copy(src, dest)
    except OSError as e:
      logging.error('Failed to copy %s to %s', src, dest)
      raise

  # Make the _ah subdirectory for the app engine tools.
  ah_dir = os.path.join(dst_deployment_dir, '_ah')
  try:
    os.mkdir(ah_dir)
  except OSError as e:
    logging.error('Failed to create %s: %s', ah_dir, e.strerror)
    raise

  # Copy gab.
  try:
    gab_dest = os.path.join(ah_dir, 'gab')
    shutil.copy(go_app_builder, gab_dest)
  except OSError as e:
    logging.error('Failed to copy %s to %s', go_app_builder, gab_dest)
    raise

  # Write build script.
  gab_args = [
      '/app/_ah/gab',
      '-app_base', '/app',
      '-arch', '6',
      '-dynamic',
      '-goroot', '/goroot',
      '-nobuild_files', '^' + str(nobuild_files),
      '-unsafe',
      '-binary_name', '_ah_exe',
      '-work_dir', '/tmp/work',
      '-vm',
  ]
  gab_args.extend(go_application.list_go_files(
      application_dir, nobuild_files, skip_files))
  gab_args.extend([x[0] for x in extras])
  dst_build = os.path.join(ah_dir, 'build.sh')
  lines = [
      '#!/bin/bash',
      'set -e',
      'mkdir -p /tmp/work',
      'chmod a+x /app/_ah/gab',
      # Without this line, Windows errors "text file busy".
      'shasum /app/_ah/gab',
      ' '.join(gab_args),
      'mv /tmp/work/_ah_exe /app/_ah/exe',
      'rm -rf /tmp/work',
      'echo Done.',
  ]
  with open(dst_build, 'wb') as fd:
    fd.write('\n'.join(lines) + '\n')
  os.chmod(dst_build, 0777)

  # TODO: Remove this when classic Go SDK is gone.
  # Write default Dockerfile if none found.
  _write_dockerfile(dst_deployment_dir)
  # Also write the default Dockerfile if not found in the app dir.
  _write_dockerfile(application_dir)
