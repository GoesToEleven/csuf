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
"""Manages a VM Runtime process running inside of a docker container."""

import logging
import os
import socket

import google

from google.appengine.tools.devappserver2 import application_configuration
from google.appengine.tools.devappserver2 import http_proxy
from google.appengine.tools.devappserver2 import http_runtime_constants
from google.appengine.tools.devappserver2 import instance
from google.appengine.tools.devappserver2 import log_manager
from google.appengine.tools.docker import containers


_APP_ENGINE_PREFIX = 'google.appengine'

# Number of seconds after a container start before we check if there is
# any old containers and images to cleanup.
_CLEANUP_DELAY_SEC = 10.0

_DOCKER_IMAGE_NAME_FORMAT = '{display}.{module}.{version}'

# This is the number of containers the cleanup process will leave on docker for
# investigation purposes.
_CONTAINERS_TO_KEEP = 10


class Error(Exception):
  """Base class for errors in this module."""


class InvalidEnvVariableError(Error):
  """Raised if an environment variable name or value cannot be supported."""


class VersionError(Error):
  """Raised if no version is specified in application configuration file."""


class InvalidForwardedPortError(Error):
  """Raised if the forwarded port is already used (for example by debugger)."""


def _GetPortToPublish(port):
  """Checks if given port is available.

  Useful for publishing debug ports when it's more convenient to bind to
  the same address on each container restart.

  Args:
    port: int, Port to check.

  Returns:
    given port if it is available, None if it is already taken (then any
        random available port will be selected by Dockerd).
  """
  sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
  try:
    sock.bind(('', port))
    sock.close()
    return port
  except socket.error:
    logging.warning('Requested debug port %d is already in use. '
                    'Will use another available port.', port)
  return None


def _ContainerName(image_name, instance_id, timestamp=None):
  """Generates a container name.

  Args:
    image_name: the base image name.
    instance_id: the instance # of the module.
    timestamp: the timestamp as a string you want to generate the name from.
               If None, it will be the now in the UTC timezone in ISO 8601.

  Returns:
    the generated container name.
  """
  return containers.CleanableContainerName(_APP_ENGINE_PREFIX,
                                           base_name='%s.%s' % (image_name,
                                                                instance_id),
                                           timestamp=timestamp)


class VMRuntimeProxy(instance.RuntimeProxy):
  """Manages a VM Runtime process running inside of a docker container."""

  DEFAULT_DEBUG_PORT = 5005

  def __init__(self, docker_client, runtime_config_getter,
               module_configuration,
               default_port=8080, port_bindings=None,
               additional_environment=None):
    """Initializer for VMRuntimeProxy.

    Args:
      docker_client: docker.Client object to communicate with Docker daemon.
      runtime_config_getter: A function that can be called without arguments
          and returns the runtime_config_pb2.Config containing the configuration
          for the runtime.
      module_configuration: An application_configuration.ModuleConfiguration
          instance respresenting the configuration of the module that owns the
          runtime.
      default_port: int, main port inside of the container that instance is
          listening on.
      port_bindings: dict, Additional port bindings from the container.
      additional_environment: doct, Additional environment variables to pass
          to the container.
    """
    super(VMRuntimeProxy, self).__init__()

    self._runtime_config_getter = runtime_config_getter
    self._module_configuration = module_configuration
    self._docker_client = docker_client
    self._default_port = default_port
    self._port_bindings = port_bindings
    self._additional_environment = additional_environment
    self._log_manager = log_manager.get(
        self._docker_client,
        enable_logging=self._runtime_config_getter().vm_config.enable_logs)
    self._container = None
    self._proxy = None

  def handle(self, environ, start_response, url_map, match, request_id,
             request_type):
    """Serves request by forwarding it to application instance via HttpProxy.

    Args:
      environ: An environ dict for the request as defined in PEP-333.
      start_response: A function with semantics defined in PEP-333.
      url_map: An appinfo.URLMap instance containing the configuration for the
          handler matching this request.
      match: A re.MatchObject containing the result of the matched URL pattern.
      request_id: A unique string id associated with the request.
      request_type: The type of the request. See instance.*_REQUEST module
          constants.

    Returns:
      Generator of sequence of strings containing the body of the HTTP response.

    Raises:
      InvalidEnvVariableError: if user tried to redefine any of the reserved
          environment variables.
    """
    return self._proxy.handle(environ, start_response, url_map, match,
                              request_id, request_type)

  def _instance_died_unexpectedly(self):
    # TODO: Check if container is still up and running
    return False

  def get_instance_logs(self):
    # TODO: Handle docker container's logs
    return ''

  def _escape_domain(self, application_external_name):
    return application_external_name.replace(':', '.')

  def start(self, dockerfile_dir=None, request_id_header_name=None):
    runtime_config = self._runtime_config_getter()

    if not self._module_configuration.major_version:
      logging.error('Version needs to be specified in your application '
                    'configuration file.')
      raise VersionError()

    self._log_manager.add(
        self._escape_domain(
            self._module_configuration.application_external_name),
        self._module_configuration.module_name,
        self._module_configuration.major_version, runtime_config.instance_id)

    if not dockerfile_dir:
      dockerfile_dir = self._module_configuration.application_root

    # api_host set to 'localhost' won't be accessible from a docker container
    # because container will have it's own 'localhost'.
    # 10.0.2.2 is a special network setup by virtualbox to connect from the
    # guest to the host.
    api_host = runtime_config.api_host
    if runtime_config.api_host in ('0.0.0.0', 'localhost'):
      api_host = '10.0.2.2'

    image_name = _DOCKER_IMAGE_NAME_FORMAT.format(
        # Escape domain if it is present.
        display=self._escape_domain(
            self._module_configuration.application_external_name),
        module=self._module_configuration.module_name,
        version=self._module_configuration.major_version)

    port_bindings = self._port_bindings if self._port_bindings else {}
    port_bindings.setdefault(self._default_port, None)
    debug_port = None

    environment = {
        'API_HOST': api_host,
        'API_PORT': runtime_config.api_port,
        'GAE_LONG_APP_ID': self._module_configuration.application_external_name,
        'GAE_PARTITION': self._module_configuration.partition,
        'GAE_MODULE_NAME': self._module_configuration.module_name,
        'GAE_MODULE_VERSION': self._module_configuration.major_version,
        'GAE_MINOR_VERSION': self._module_configuration.minor_version,
        'GAE_MODULE_INSTANCE': runtime_config.instance_id,
        'GAE_SERVER_PORT': runtime_config.server_port,
        'MODULE_YAML_PATH': os.path.basename(
            self._module_configuration.config_path),
        'SERVER_SOFTWARE': http_runtime_constants.SERVER_SOFTWARE
    }
    if self._additional_environment:
      environment.update(self._additional_environment)

    # Handle user defined environment variables
    if self._module_configuration.env_variables:
      ev = (environment.viewkeys() &
            self._module_configuration.env_variables.viewkeys())
      if ev:
        raise InvalidEnvVariableError(
            'Environment variables [%s] are reserved for App Engine use' %
            ', '.join(ev))

      environment.update(self._module_configuration.env_variables)

      # Publish debug port if running in Debug mode.
      if self._module_configuration.env_variables.get('DBG_ENABLE'):
        debug_port = int(self._module_configuration.env_variables.get(
            'DBG_PORT', self.DEFAULT_DEBUG_PORT))
        environment['DBG_PORT'] = debug_port
        port_bindings[debug_port] = _GetPortToPublish(debug_port)

    # Publish forwarded ports
    # NOTE: fowarded ports are mapped as host_port => container_port,
    # port_bindings are mapped the other way around.
    for h, c in self._module_configuration.forwarded_ports.iteritems():
      if c in port_bindings:
        raise InvalidForwardedPortError(
            'Port {port} is already used by debugger or runtime specific '
            'VM Service. Please use a different forwarded_port.'.format(port=c))
      port_bindings[c] = h

    external_logs_path = os.path.join(
        '/var/log/app_engine',
        self._escape_domain(
            self._module_configuration.application_external_name),
        self._module_configuration.module_name,
        self._module_configuration.major_version,
        runtime_config.instance_id)
    internal_logs_path = '/var/log/app_engine'
    container_name = _ContainerName(
        image_name=image_name,
        instance_id=runtime_config.instance_id)
    self._container = containers.Container(
        self._docker_client,
        containers.ContainerOptions(
            image_opts=containers.ImageOptions(
                dockerfile_dir=dockerfile_dir,
                tag=image_name,
                nocache=False),
            port=self._default_port,
            port_bindings=port_bindings,
            environment=environment,
            volumes={
                external_logs_path: {'bind': internal_logs_path}
            },
            name=container_name
        ))

    self._container.Start()
    # As we add stuff, asynchronously check later for a cleanup.
    containers.StartDelayedCleanup(
        self._docker_client, _APP_ENGINE_PREFIX, _CLEANUP_DELAY_SEC,
        _CONTAINERS_TO_KEEP)

    # Print the debug information before connecting to the container
    # as debugging might break the runtime during initialization, and
    # connecting the debugger is required to start processing requests.
    if debug_port:
      logging.info('To debug module {module} attach to {host}:{port}'.format(
          module=self._module_configuration.module_name,
          host=self.ContainerHost(),
          port=self.PortBinding(debug_port)))

    self._proxy = http_proxy.HttpProxy(
        host=self._container.host, port=self._container.port,
        instance_died_unexpectedly=self._instance_died_unexpectedly,
        instance_logs_getter=self.get_instance_logs,
        error_handler_file=application_configuration.get_app_error_file(
            self._module_configuration),
        request_id_header_name=request_id_header_name)

    # If forwarded ports are used we do not really have to serve on 8080.
    # We'll log /_ah/start request fail, but with health-checks disabled
    # we should ignore that and continue working (accepting requests on
    # our forwarded ports outside of dev server control).
    health_check = self._module_configuration.health_check
    health_check_enabled = health_check and health_check.enable_health_check

    if health_check_enabled or not self._module_configuration.forwarded_ports:
      self._proxy.wait_for_connection()

  def quit(self):
    """Kills running container and removes it."""
    self._container.Stop()

  def PortBinding(self, port):
    """Get the host binding of a container port.

    Args:
      port: Port inside container.

    Returns:
      Port on the host system mapped to the given port inside of
          the container.
    """
    return self._container.PortBinding(port)

  def ContainerHost(self):
    """Get the host IP address of the container.

    Returns:
      IP address on the host system for accessing the container.
    """
    return self._container.host
