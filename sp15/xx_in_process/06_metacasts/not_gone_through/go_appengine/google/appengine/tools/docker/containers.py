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
"""Docker image and docker container classes.

In Docker terminology image is a read-only layer that never changes.
Container is created once you start a process in Docker from an Image. Container
consists of read-write layer, plus information about the parent Image, plus
some additional information like its unique ID, networking configuration,
and resource limits.
For more information refer to http://docs.docker.io/.

Mapping to Docker CLI:
Image is a result of "docker build path/to/Dockerfile" command.
Container is a result of "docker run image_tag" command.
ImageOptions and ContainerOptions allow to pass parameters to these commands.

Versions 1.9 and 1.10 of docker remote API are supported.
"""

from collections import namedtuple
import datetime
import json
import logging
import os
import re
import socket
import ssl
import sys
import threading
import urlparse

import google
import docker
import requests


_SUCCESSFUL_BUILD_PATTERN = re.compile(r'Successfully built ([a-zA-Z0-9]{12})')

_STREAM = 'stream'


_cleanup_scheduled = None
_cleanup_scheduled_lock = threading.Lock()
_cleanup_lock = threading.Lock()

DEFAULT_LINUX_DOCKER_HOST = '/var/run/docker.sock'
DOCKER_CONNECTION_ERROR = (
    'Couldn\'t connect to the docker daemon.\n'
    'Please check if the environment variables DOCKER_HOST, '
    'DOCKER_CERT_PATH and DOCKER_TLS_VERIFY are set correctly. '
    'If you are using boot2docker, you can set them up by '
    'executing the commands that are shown by:\n'
    'boot2docker shellinit')
# TODO: revisit this message when
# https://github.com/boot2docker/boot2docker-cli/issues/301 is fixed

# This is the recognized format for container names to be automatically
# cleaned up by the _CleanupOldContainersAndImagesWithPrefix method below.
_CONTAINER_NAME_FORMAT = '{prefix}.{base_name}.{timestamp}'


def CleanableContainerName(prefix, base_name, timestamp=None):
  """Generates a container name tagged for later deletion.

  We have an asynchronous cleaning process in the SDK that can recognize these
  containers by their names and periodically cleans them up.
  The naming scheme avoids cleaning up containers that did not come
  from the SDK.

  Args:
    prefix: the prefix used to identify the cleanable containers.
    base_name: the identifier of the series of containers.
    timestamp: the timestamp as a string you want to generate the name from.
               If None, it will be the now in the UTC timezone in ISO 8601.

  Returns:
    the generated container name.
  """
  if timestamp is None:
    # ":" is not allowed in container names but is optional in
    # ISO8601.
    timestamp = datetime.datetime.utcnow().strftime('%Y-%m-%dT%H%M%S.%fZ')

  return _CONTAINER_NAME_FORMAT.format(
      prefix=prefix, base_name=base_name, timestamp=timestamp)


class ImageOptions(namedtuple('ImageOptionsT',
                              ['dockerfile_dir', 'tag', 'nocache', 'rm'])):
  """Options for building Docker Images."""

  def __new__(cls, dockerfile_dir=None, tag=None, nocache=False, rm=True):
    """This method is redefined to provide default values for namedtuple.

    Args:
      dockerfile_dir: str, Path to the directory with the Dockerfile. If it is
          None, no build is needed. We will be looking for the existing image
          with the specified tag and raise an error if it does not exist.
      tag: str, Repository name (and optionally a tag) to be applied to the
          image in case of successful build. If dockerfile_dir is None, tag
          is used for lookup of an image.
      nocache: boolean, True if cache should not be used when building the
          image.
      rm: boolean, True if intermediate images should be removed after a
          successful build. Default value is set to True because this is the
          default value used by "docker build" command.

    Returns:
      ImageOptions object.
    """
    return super(ImageOptions, cls).__new__(
        cls, dockerfile_dir=dockerfile_dir, tag=tag, nocache=nocache, rm=rm)


class ContainerOptions(namedtuple('ContainerOptionsT',
                                  ['image_opts', 'port', 'port_bindings',
                                   'environment', 'volumes', 'volumes_from',
                                   'links', 'name', 'command'])):
  """Options for creating and running Docker Containers."""

  def __new__(cls, image_opts=None, port=None, port_bindings=None,
              environment=None, volumes=None, volumes_from=None, links=None,
              name=None, command=None):
    """This method is redefined to provide default values for namedtuple.

    Args:
      image_opts: ImageOptions, properties of underlying Docker Image.
      port: int, Primary port that the process inside of a container is
          listening on. If this port is not part of the port bindings
          specified, a default binding will be added for this port.
      port_bindings: dict, Port bindings for exposing multiple ports. If the
          only binding needed is the default binding of just one port this
          can be None.
      environment: dict, Environment variables.
      volumes: dict,  Volumes to mount from the host system.
      volumes_from: list, Volumes from the specified container(s).
      links: dict, Links to the specified container(s).
      name: str, Name of a container. Needed for data containers.
      command: str, The command to execute within the container.

    Returns:
      ContainerOptions object.
    """
    return super(ContainerOptions, cls).__new__(
        cls, image_opts=image_opts, port=port, port_bindings=port_bindings,
        environment=environment, volumes=volumes, volumes_from=volumes_from,
        name=name, links=links, command=command)


class Error(Exception):
  """Base exception for containers module."""


class ImageError(Error):
  """Image related errors."""


class ContainerError(Error):
  """Container related erorrs."""


class DockerDaemonConnectionError(Error):
  """Raised if the docker client can't connect to the docker daemon."""


class BaseImage(object):
  """Abstract base class for Docker images."""

  def __init__(self, docker_client, image_opts):
    """Initializer for BaseImage.

    Args:
      docker_client: an object of docker.Client class to communicate with a
          Docker daemon.
      image_opts: an instance of ImageOptions class describing the parameters
          passed to docker commands.
    """
    self._docker_client = docker_client
    self._image_opts = image_opts
    self._id = None

  def Build(self):
    """Calls "docker build" if needed."""
    raise NotImplementedError

  def Remove(self):
    """Calls "docker rmi" if needed."""
    raise NotImplementedError

  @property
  def id(self):
    """Returns 64 hexadecimal digit string identifying the image."""
    # Might also be a first 12-characters shortcut.
    return self._id

  @property
  def tag(self):
    """Returns image tag string."""
    return self._image_opts.tag

  def __enter__(self):
    """Makes BaseImage usable with "with" statement."""
    self.Build()
    return self

  # pylint: disable=redefined-builtin
  def __exit__(self, type, value, traceback):
    """Makes BaseImage usable with "with" statement."""
    self.Remove()

  def __del__(self):
    """Makes sure that build artifacts are cleaned up."""
    self.Remove()


class Image(BaseImage):
  """Docker image that requires building and should be removed afterwards."""

  def __init__(self, docker_client, image_opts):
    """Initializer for Image.

    Args:
      docker_client: an object of docker.Client class to communicate with a
          Docker daemon.
      image_opts: an instance of ImageOptions class that must have
          dockerfile_dir set. image_id will be returned by "docker build"
          command.

    Raises:
      ImageError: if dockerfile_dir is not set.
    """
    if not image_opts.dockerfile_dir:
      raise ImageError('dockerfile_dir for images that require building '
                       'must be set.')

    super(Image, self).__init__(docker_client, image_opts)

  def Build(self):
    """Calls "docker build".

    Raises:
      ImageError: if the image could not be built.
    """
    logging.info('Building docker image %s from %s/Dockerfile:',
                 self.tag, self._image_opts.dockerfile_dir)

    logging.info('-' * 20 + '  DOCKER BUILD  ' + '-' * 20)

    build_res = self._docker_client.build(
        path=self._image_opts.dockerfile_dir,
        tag=self.tag,
        quiet=False, fileobj=None, nocache=self._image_opts.nocache,
        rm=self._image_opts.rm)

    error = None
    error_detail = None
    log_records = []
    try:
      for line in build_res:
        line = line.strip()
        if not line:
          continue
        log_record = json.loads(line)
        log_records.append(log_record)
        if 'stream' in log_record:
          logging.info(log_record['stream'].strip())
        elif 'error' in log_record:
          error = log_record['error'].strip()
          logging.error(error)
        elif 'errorDetail' in log_record:
          error_detail = log_record['errorDetail']['message'].strip()
          logging.error('Details: ' + error_detail)
    except docker.errors.APIError as e:
      logging.error(e.explanation)
      error = e.explanation
      error_detail = ''
    finally:
      logging.info('-' * 56)

    if not log_records:
      raise ImageError('Error building docker image %s [with no output]',
                       self.tag)

    success_message = log_records[-1].get(_STREAM)
    if success_message:
      m = _SUCCESSFUL_BUILD_PATTERN.match(success_message)
      if m:
        # The build was successful.
        self._id = m.group(1)
        logging.info('Image %s built, id = %s', self.tag, self.id)
        return
    msg = 'Docker build aborted: %s' % error
    if error_detail:
      msg += '\nMore info: %s' % error_detail
    raise ImageError(msg)

  def Remove(self):
    """Calls "docker rmi"."""
    # This will be done automatically by the cleanup.
    self._id = None


class PrebuiltImage(BaseImage):
  """Prebuilt Docker image. Build and Remove functions are noops."""

  def __init__(self, docker_client, image_opts):
    """Initializer for PrebuiltImage.

    Args:
      docker_client: an object of docker.Client class to communicate with a
          Docker daemon.
      image_opts: an instance of ImageOptions class that must have
          dockerfile_dir not set and tag set.

    Raises:
      ImageError: if image_opts.dockerfile_dir is set or
          image_opts.tag is not set.
    """
    if image_opts.dockerfile_dir:
      raise ImageError('dockerfile_dir for PrebuiltImage must not be set.')

    if not image_opts.tag:
      raise ImageError('PrebuiltImage must have tag specified to find '
                       'image id.')

    super(PrebuiltImage, self).__init__(docker_client, image_opts)

  def Build(self):
    """Searches for pre-built image with specified tag.

    Raises:
      ImageError: if image with this tag was not found.
    """
    logging.info('Looking for image_id for image with tag %s', self.tag)
    images = self._docker_client.images(
        name=self.tag, quiet=True, all=False, viz=False)

    if not images:
      raise ImageError('Image with tag %s was not found' % self.tag)

    # TODO: check if it's possible to have more than one image returned.
    self._id = images[0]

  def Remove(self):
    """Unassigns image_id only, does not remove the image as we don't own it."""
    self._id = None


def CreateImage(docker_client, image_opts):
  """Creates an new object to represent Docker image.

  Args:
    docker_client: an object of docker.Client class to communicate with a
        Docker daemon.
    image_opts: an instance of ImageOptions class.

  Returns:
    New object, subclass of BaseImage class.
  """
  image = Image if image_opts.dockerfile_dir else PrebuiltImage
  return image(docker_client, image_opts)


def GetDockerHost(docker_client):
  parsed_url = urlparse.urlparse(docker_client.base_url)

  # Socket url schemes look like: unix:// or http+unix://.
  # If the user is running docker locally and connecting over a socket, we
  # should just use localhost.
  if 'unix' in parsed_url.scheme:
    return 'localhost'
  return parsed_url.hostname


def _GetAllLingeringContainersInfo(docker_client, prefix):
  """Lists all the stopped App engine containers.

  Args:
    docker_client: an object of docker.Client class to communicate with a
        Docker daemon.
    prefix: str, the container name prefix we are looking for.

  Returns:
    A list of container_info dictionaries.
  """
  all_containers = docker_client.containers(all=True)
  # We roundtrip to the client as 'Status' is only a human readable string.
  live_containers_ids = [container_info['Id']
                         for container_info in docker_client.containers(
                             quiet=True, all=False)]

  def IsPrefixedAndStopped(cinfo):
    return any(
        name.startswith('/' + prefix)
        for name in cinfo['Names']) and cinfo['Id'] not in live_containers_ids

  return filter(IsPrefixedAndStopped, all_containers)


def StartDelayedCleanup(docker_client, prefix, delay_sec,
                        old_instances_to_spare=0):
  """Start later a cleanup of the stopped containers with a prefix.

  Args:
    docker_client: an object of docker.Client class to communicate with a
      Docker daemon.
    prefix: str, the container name prefix we want to cleanup.
    delay_sec: The delay before we trigger it.
    old_instances_to_spare: leave at least this amount of old containers.
  """
  global _cleanup_scheduled
  with _cleanup_scheduled_lock:
    if not _cleanup_scheduled:
      _cleanup_scheduled = threading.Timer(
          delay_sec, _CleanupOldContainersAndImagesWithPrefix,
          [docker_client, prefix, old_instances_to_spare])
      _cleanup_scheduled.daemon = True
      _cleanup_scheduled.start()


def _CleanupOldContainersAndImagesWithPrefix(docker_client, prefix,
                                             containers_to_keep=0):
  """Remove Old App Engine unused containers and images.

  Args:
    docker_client: an object of docker.Client class to communicate with a
        Docker daemon.
    prefix: str, the container name prefix we are looking for.
    containers_to_keep: leave at least this amount of old containers.
  """
  global _cleanup_scheduled
  # Directly at the beginning of the cleanup we can authorize the scheduling of
  # another cleanup.
  with _cleanup_scheduled_lock:
    _cleanup_scheduled = None

  with _cleanup_lock:
    logging.debug('Automatic cleanup...')
    stopped_containers = _GetAllLingeringContainersInfo(
        docker_client, prefix)
    by_creation = lambda container_info: container_info['Created']
    stopped_containers.sort(key=by_creation)
    to_remove = stopped_containers[:-containers_to_keep]
    for container in to_remove:
      try:
        logging.debug('Removing old container: %s:%s',
                      container['Id'], container['Names'])
        docker_client.remove_container(container)
      except docker.errors.APIError as e:
        logging.warning('Container %s (id=%s) cannot be removed: %s.\n'
                        'Try cleaning up old containers manually.\n'
                        'They can be listed with "docker ps -a".',
                        container['Names'], container['Id'], e)
      else:
        try:
          img = container['Image']
          if [cinfo for cinfo in docker_client.containers(all=True)
              if cinfo['Image'] == img]:
            # another container derived from this image, clean it up later.
            continue
          logging.debug('Removing old image: %s', img)

          # TODO: remove once it is fixed upstream.
          # see https://github.com/docker/docker/issues/9602
          if isinstance(img, (int, long)):
            logging.error('Some unnamed images could not be removed.\n'
                          'Try cleaning up old images manually.\n'
                          'They can be listed with "docker images".')
            continue
          docker_client.remove_image(img)
        except docker.errors.APIError as e:
          logging.warning('Image Id %s cannot be removed: %s.\n'
                          'Try cleaning up old images manually.\n'
                          'They can be listed with "docker images".',
                          container['Image'], e)
    logging.debug('Cleanup finished.')


class Container(object):
  """Docker Container."""

  def __init__(self, docker_client, container_opts):
    """Initializer for Container.

    Args:
      docker_client: an object of docker.Client class to communicate with a
          Docker daemon.
      container_opts: an instance of ContainerOptions class.
    """
    self._docker_client = docker_client
    self._container_opts = container_opts

    self._image = CreateImage(docker_client, container_opts.image_opts)
    self._id = None
    self._host = GetDockerHost(self._docker_client)
    self._container_host = None
    # Port bindings will be set to a dictionary mapping exposed ports
    # to the interface they are bound to. This will be populated from
    # the container options passed when the container is started.
    self._port_bindings = None

    # Use the daemon flag in case we leak these threads.
    self._logs_listener = threading.Thread(target=self._ListenToLogs)
    self._logs_listener.daemon = True

  def Start(self):
    """Builds an image (if necessary) and runs a container.

    Raises:
      ContainerError: if container_id is already set, i.e. container is already
          started or we could not build the underlying image.
    """
    if self.id:
      raise ContainerError('Trying to start already running container.')

    # don't concurrently create and cleanup
    with _cleanup_lock:
      self._image.Build()

      logging.info('Creating container...')
      port_bindings = self._container_opts.port_bindings or {}
      if self._container_opts.port:
        # Add primary port to port bindings if not already specified.
        # Setting its value to None lets docker pick any available port.
        port_bindings[self._container_opts.port] = port_bindings.get(
            self._container_opts.port)

      response = self._docker_client.create_container(
          image=self._image.id, hostname=None, user=None, detach=True,
          command=self._container_opts.command,
          stdin_open=False,
          tty=False, mem_limit=0,
          ports=port_bindings.keys(),
          volumes=(self._container_opts.volumes.keys()
                   if self._container_opts.volumes else None),
          environment=self._container_opts.environment,
          dns=None,
          network_disabled=False,
          name=self.name)

      self._id = response.get('Id')
      warnings = response.get('Warnings')
      if warnings:
        logging.warning(warnings)

      logging.info('Container %s created.', self.id)

      self._docker_client.start(
          self.id,
          port_bindings=port_bindings,
          binds=self._container_opts.volumes,
          links=self._container_opts.links,
          # In the newer API version volumes_from got moved from
          # create_container to start. In older version volumes_from option was
          # completely broken therefore we support only passing volumes_from
          # in start.
          volumes_from=self._container_opts.volumes_from)

      self._logs_listener.start()

      if not port_bindings:
        # Nothing to inspect
        return

      container_info = self._docker_client.inspect_container(self._id)
      network_settings = container_info['NetworkSettings']
      self._container_host = network_settings['IPAddress']
      self._port_bindings = {
          port: int(network_settings['Ports']['%d/tcp' % port][0]['HostPort'])
          for port in port_bindings
      }

  def Stop(self):
    """Stops a running container, removes it and underlying image if needed."""
    if self._id:
      self._docker_client.kill(self.id)
      self._id = None

  def PortBinding(self, port):
    """Get the host binding of a container port.

    Args:
      port: Port inside container.

    Returns:
      Port on the host system mapped to the given port inside of
          the container.
    """
    return self._port_bindings.get(port)

  @property
  def host(self):
    """Host the container can be reached at by the host (i.e. client) system."""
    return self._host

  @property
  def port(self):
    """Port (on the host system) mapped to the port inside of the container."""
    return self._port_bindings[self._container_opts.port]

  @property
  def addr(self):
    """An address the container can be reached at by the host system."""
    return '%s:%d' % (self.host, self.port)

  @property
  def id(self):
    """Returns 64 hexadecimal digit string identifying the container."""
    return self._id

  @property
  def container_addr(self):
    """An address the container can be reached at by another container."""
    return '%s:%d' % (self._container_host, self._container_opts.port)

  @property
  def name(self):
    """String, identifying a container. Required for data containers."""
    return self._container_opts.name

  def _ListenToLogs(self):
    """Logs all output from the docker container.

    The docker.Client.logs method returns a generator that yields log lines.
    This method iterates over that generator and outputs those log lines to
    the devappserver2 logs.
    """
    log_lines = self._docker_client.logs(container=self.id, stream=True)
    for line in log_lines:
      line = line.strip()
      logging.debug('Container: %s: %s', self.id[0:12], line)

  def __enter__(self):
    """Makes Container usable with "with" statement."""
    self.Start()
    return self

  # pylint: disable=redefined-builtin
  def __exit__(self, type, value, traceback):
    """Makes Container usable with "with" statement."""
    self.Stop()

  def __del__(self):
    """Makes sure that all build and run artifacts are cleaned up."""
    self.Stop()


def _KwargsFromEnv():
  """Helper to build docker.Client constructor kwargs from the environment."""
  host = os.environ.get('DOCKER_HOST')
  cert_path = os.environ.get('DOCKER_CERT_PATH')
  tls_verify = os.environ.get('DOCKER_TLS_VERIFY')
  logging.debug('Detected docker environment variables: DOCKER_HOST=%s, '
                'DOCKER_CERT_PATH=%s, DOCKER_TLS_VERIFY=%s', host, cert_path,
                tls_verify)

  params = {}

  if host:
    params['base_url'] = (host.replace('tcp://', 'https://') if tls_verify
                          else host)
  elif sys.platform.startswith('linux'):
    # if this is a linux user, the default value of DOCKER_HOST should be the
    # unix socket.

    # first check if the socket is valid to give a better feedback to the user.
    if os.path.exists(DEFAULT_LINUX_DOCKER_HOST):
      sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
      try:
        sock.connect(DEFAULT_LINUX_DOCKER_HOST)
        params['base_url'] = 'unix://' + DEFAULT_LINUX_DOCKER_HOST
      except socket.error:
        logging.warning('Found a stale /var/run/docker.sock, '
                        'did you forget to start your docker daemon?')
      finally:
        sock.close()

  if tls_verify and cert_path:
    # assert_hostname=False is needed for boot2docker to work with our custom
    # registry.
    params['tls'] = docker.docker.tls.TLSConfig(
        client_cert=(os.path.join(cert_path, 'cert.pem'),
                     os.path.join(cert_path, 'key.pem')),
        ca_cert=os.path.join(cert_path, 'ca.pem'),
        verify=True,
        ssl_version=ssl.PROTOCOL_TLSv1,
        assert_hostname=False)
  return params


def NewDockerClientNoCheck(**kwargs):
  """Factory method for building a docker.Client from environment variables.

  Args:
    **kwargs: Any kwargs will be passed to the docker.Client constructor and
      override any determined from the environment.

  Returns:
    A docker.Client instance.

  Raises:
    DockerDaemonConnectionError: If the docker daemon isn't responding.
  """
  kwargs_from_env = _KwargsFromEnv()
  kwargs_from_env.update(kwargs)
  if 'base_url' not in kwargs_from_env:
    raise DockerDaemonConnectionError(DOCKER_CONNECTION_ERROR)
  return docker.Client(**kwargs_from_env)


def NewDockerClient(**kwargs):
  """Factory method for building a docker.Client from environment variables.

  Args:
    **kwargs: Any kwargs will be passed to the docker.Client constructor and
      override any determined from the environment.

  Returns:
    A docker.Client instance.

  Raises:
    DockerDaemonConnectionError: If the docker daemon isn't responding.
  """
  client = NewDockerClientNoCheck(**kwargs)
  try:
    client.ping()
  except requests.exceptions.ConnectionError, e:
    logging.error('Failed to connect to Docker Daemon due to: %s', e)
    raise DockerDaemonConnectionError(DOCKER_CONNECTION_ERROR)
  return client
