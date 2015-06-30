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
"""Manages creation of VM Runtime instances."""

import logging

import google

from google.appengine.api import appinfo
from google.appengine.tools.devappserver2 import instance
from google.appengine.tools.devappserver2 import vm_runtime_proxy
from google.appengine.tools.devappserver2 import vm_runtime_proxy_go
from google.appengine.tools.docker import containers


# docker uses requests which logs lots of noise.
logging.getLogger('requests').setLevel(logging.WARNING)


class VMRuntimeInstanceFactory(instance.InstanceFactory):
  """A factory that creates new VM runtime Instances."""

  START_URL_MAP = appinfo.URLMap(
      url='/_ah/start',
      script='/dev/null',
      login='admin')
  WARMUP_URL_MAP = appinfo.URLMap(
      url='/_ah/warmup',
      script='/dev/null',
      login='admin')

  RUNTIME_SPECIFIC_PROXY = {
      'go': vm_runtime_proxy_go.GoVMRuntimeProxy,
  }

  SUPPORTS_INTERACTIVE_REQUESTS = True
  FILE_CHANGE_INSTANCE_RESTART_POLICY = instance.ALWAYS

  # Timeout of HTTP request from docker-py client to docker daemon, in seconds.
  DOCKER_D_REQUEST_TIMEOUT_SECS = 60

  def __init__(self, request_data, runtime_config_getter, module_configuration):
    """Initializer for VMRuntimeInstanceFactory.

    Args:
      request_data: A wsgi_request_info.WSGIRequestInfo that will be provided
          with request information for use by API stubs.
      runtime_config_getter: A function that can be called without arguments
          and returns the runtime_config_pb2.Config containing the configuration
          for the runtime.
      module_configuration: An application_configuration.ModuleConfiguration
          instance representing the configuration of the module that owns the
          runtime.
    """
    super(VMRuntimeInstanceFactory, self).__init__(
        request_data,
        8 if runtime_config_getter().threadsafe else 1, 10)
    self._runtime_config_getter = runtime_config_getter
    self._module_configuration = module_configuration
    self._docker_client = containers.NewDockerClient(
        version='1.9',
        timeout=self.DOCKER_D_REQUEST_TIMEOUT_SECS)

  def new_instance(self, instance_id, expect_ready_request=False):
    """Create and return a new Instance.

    Args:
      instance_id: A string or integer representing the unique (per module) id
          of the instance.
      expect_ready_request: If True then the instance will be sent a special
          request (i.e. /_ah/warmup or /_ah/start) before it can handle external
          requests.

    Returns:
      The newly created instance.Instance.
    """

    def runtime_config_getter():
      runtime_config = self._runtime_config_getter()
      runtime_config.instance_id = str(instance_id)
      return runtime_config

    effective_runtime = self._module_configuration.effective_runtime
    proxy_class = self.RUNTIME_SPECIFIC_PROXY.get(
        effective_runtime, vm_runtime_proxy.VMRuntimeProxy)

    proxy = proxy_class(
        self._docker_client, runtime_config_getter, self._module_configuration)
    return instance.Instance(self.request_data,
                             instance_id,
                             proxy,
                             self.max_concurrent_requests,
                             self.max_background_threads,
                             expect_ready_request)
