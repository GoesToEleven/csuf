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
"""LogManager for Managed VMs modules.

Should be accessed by get() function.
"""

import atexit
import httplib
import logging
import os
import threading
import urllib

import google

from google.pyglib import singleton

from google.appengine.tools.devappserver2 import http_utils
from google.appengine.tools.docker import containers


APP_ENGINE_LOG_SERVER_HOST = 'APP_ENGINE_LOG_SERVER_HOST'
APP_ENGINE_LOG_SERVER_PORT = 'APP_ENGINE_LOG_SERVER_PORT'

_DB_PATH = '/var/log/sqlite'
_LOGS_PATH = '/var/log/app_engine'
_TD_AGENT_PATH = '/var/tmp/td-agent'

_LOG_PROCESSOR_IMAGE = 'google/appengine-log-processor'
_LOG_SERVER_IMAGE = 'google/appengine-log-server'
_DEFAULT_LOG_SERVER_PORT = 8080

_LOG_TYPES = ['app', 'appjson', 'request']

_APP_ENGINE_PREFIX = 'google.appengine'


# TODO: more escaping.
def _escape(s):
  return s.replace('-', '_')


def _make_container_name(app, module, version, instance):
  base_name_tmpl = '{app}.{module}.{version}.{instance}.logs'
  base_name = _escape(base_name_tmpl.format(app=app, module=module,
                                            version=version, instance=instance))
  return containers.CleanableContainerName(_APP_ENGINE_PREFIX, base_name)


def _make_external_logs_path(app, module, version, instance):
  return os.path.join(_LOGS_PATH,
                      app, module, version, instance)


def _describe_volume(internal, external=None):
  return (external if external else internal), {'bind': internal}


class _LogManagerDisabled(object):
  """Base class for Log Managers. Logs are disabled by default."""

  def __init__(self, docker_client, log_server_port):
    pass

  def start(self):
    pass

  def stop(self):
    pass

  def add(self, app, module, version, instance):
    pass

  @property
  def host(self):
    return ''

  @property
  def port(self):
    return -1


class _LogManager(_LogManagerDisabled):
  """Manages creation of log server and log processors for each instance."""

  def __init__(self, docker_client, log_server_port):
    super(_LogManager, self).__init__(docker_client, log_server_port)

    self._docker_client = docker_client

    volumes = [_describe_volume(_DB_PATH)]
    self._server = containers.Container(
        self._docker_client,
        containers.ContainerOptions(
            image_opts=containers.ImageOptions(tag=_LOG_SERVER_IMAGE),
            port=log_server_port,
            volumes=dict(volumes),
            name=containers.CleanableContainerName(_APP_ENGINE_PREFIX,
                                                   'log-server')))

    self._lock = threading.RLock()
    self._containers = {}

  def start(self):
    self._server.Start()
    http_utils.wait_for_connection(self._server.host, self._server.port, 100)

  def stop(self):
    for c in self._containers.itervalues():
      c.Stop()
    self._server.Stop()

  def add(self, app, module, version, instance):
    container_name = _make_container_name(app, module, version, instance)

    def _create_table(log_type):
      """Sends a request to log-server container to create a table if needed."""
      params = urllib.urlencode({
          'app': _escape(app), 'module': _escape(module),
          'version': _escape(version), 'instance': _escape(instance),
          'log_type': log_type})
      headers = {
          'Content-Type': 'application/x-www-form-urlencoded',
          'Accept': 'text/plain'}

      conn = httplib.HTTPConnection(self._server.host, self._server.port)
      conn.request('POST', '/submit', params, headers)
      response = conn.getresponse()
      logging.debug(
          'Sent table creation request to {host}:{port}?{params}. '
          'Received {status}, reason: {reason}.'.format(
              host=self._server.host, port=self._server.port,
              params=params, status=response.status, reason=response.reason))

    def _make_logs_container():
      """Creates a log-processor container."""

      environment = {
          'LOGS_PATH': _LOGS_PATH,
          'PREFIX': _escape('{app}_{module}'
                            '_{version}_{instance}'.format(app=app,
                                                           module=module,
                                                           version=version,
                                                           instance=instance))
      }

      volumes = [
          _describe_volume(_LOGS_PATH,
                           _make_external_logs_path(app, module,
                                                    version, instance)),
          _describe_volume(_DB_PATH),
          _describe_volume(_TD_AGENT_PATH)
      ]

      return containers.Container(
          self._docker_client,
          containers.ContainerOptions(
              image_opts=containers.ImageOptions(tag=_LOG_PROCESSOR_IMAGE),
              environment=environment,
              volumes=dict(volumes),
              name=container_name))
    with self._lock:
      if container_name in self._containers:
        return

      for l in _LOG_TYPES:
        _create_table(l)

      container = _make_logs_container()
      self._containers[container_name] = container
    container.Start()

  @property
  def host(self):
    return self._server.host

  @property
  def port(self):
    return self._server.port


@singleton.Singleton
class LogManagerDisabled(_LogManagerDisabled):
  """Singleton instance of _LogManagerDisabled."""


@singleton.Singleton
class LogManager(_LogManager):
  """Singleton instance of _LogManager."""


# TODO: images lookup before confirming that logs are enabled.
def get(docker_client=None, log_server_port=_DEFAULT_LOG_SERVER_PORT,
        enable_logging=False):
  """Returns a LogManager/LogManagerDisabled instance. Creates one if needed."""
  c = LogManager if enable_logging else LogManagerDisabled
  try:
    instance = c(docker_client, log_server_port)
    atexit.register(c.stop, instance)
    instance.start()

    # To pass these values to Admin Server to query logs.
    os.environ[APP_ENGINE_LOG_SERVER_HOST] = instance.host
    os.environ[APP_ENGINE_LOG_SERVER_PORT] = str(instance.port)
  except singleton.ConstructorCalledAgainError:
    instance = c.Singleton()
  return instance
