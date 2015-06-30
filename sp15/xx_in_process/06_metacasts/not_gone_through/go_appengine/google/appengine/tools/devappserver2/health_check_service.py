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
"""A health checking implementation for the SDK.

<scrub>
This code attempts to match the production implementation as closely as
possible in apphosting/runtime/vm/vm_health_check.cc.

One instance of HealthChecker should be created per instance.Instance that has
health checking enabled. The HealthChecker instance needs to be started, but
will stop itself automatically.
</scrub>
"""


import logging
import threading
import time

from google.appengine.api import request_info
from google.appengine.tools.devappserver2 import start_response_utils


class _HealthCheckState(object):
  """A class to track the state of a health checked instance."""

  def __init__(self):
    """Initializes a _HealthCheckState object."""
    self.consecutive_healthy_responses = 0
    self.consecutive_unhealthy_responses = 0
    self.is_last_successful = False

  def update(self, healthy):
    """Updates the state.

    Args:
      healthy: Bool indicating whether the last attempt was healthy.
    """
    self.is_last_successful = healthy
    if healthy:
      self.consecutive_healthy_responses += 1
      self.consecutive_unhealthy_responses = 0
    else:
      self.consecutive_healthy_responses = 0
      self.consecutive_unhealthy_responses += 1

  def __str__(self):
    """Outputs the state in a readable way for logging."""
    tmpl = '{number} consecutive {state} responses.'
    if self.consecutive_healthy_responses:
      number = self.consecutive_healthy_responses
      state = 'HEALTHY'
    else:
      number = self.consecutive_unhealthy_responses
      state = 'UNHEALTHY'
    return tmpl.format(number=number, state=state)


class HealthChecker(object):
  """A class to perform health checks for an instance.

  This class uses the settings specified in appinfo.HealthCheck and the
  callback specified to check the health of the specified instance. When
  appropriate, this class changes the state of the specified instance so it is
  placed into or taken out of load balancing. This class will also use another
  callback to restart the instance, if necessary.
  """

  def __init__(self, instance, config, send_request, restart):
    """Initializes a HealthChecker object.

    Args:
      instance: An instance.Instance object.
      config: An appinfo.HealthCheck object.
      send_request: A function to call that makes the health check request.
      restart: A function to call that restarts the instance.
    """
    self._instance = instance
    self._config = config
    self._send_request = send_request
    self._restart = restart

  def start(self):
    """Starts the health checks."""
    self._instance.set_health(False)
    logging.info('Health checks starting for instance %s.',
                 self._instance.instance_id)
    loop = threading.Thread(target=self._loop, name='Health Check')
    loop.daemon = True
    loop.start()

  def _should_continue(self):
    return self._running and not self._instance.has_quit

  def _loop(self):
    """Performs health checks and updates state over time."""
    state = _HealthCheckState()
    self._running = True
    while self._should_continue():
      logging.debug('Performing health check for instance %s.',
                    self._instance.instance_id)
      self._do_health_check(state)
      logging.debug('Health check state for instance: %s: %s',
                    self._instance.instance_id, state)
      time.sleep(self._config.check_interval_sec)

  def _do_health_check(self, state):
    health = self._get_health_check_response(state.is_last_successful)
    state.update(health)
    self._maybe_update_instance(state)

  def _maybe_update_instance(self, state):
    """Performs any required actions on the instance based on the state.

    Args:
      state: A _HealthCheckState object.
    """
    if (state.consecutive_unhealthy_responses >=
        self._config.unhealthy_threshold):
      self._instance.set_health(False)
    elif (state.consecutive_healthy_responses >=
          self._config.healthy_threshold):
      self._instance.set_health(True)

    if (state.consecutive_unhealthy_responses >=
        self._config.restart_threshold):
      self._restart_instance()

  def _get_health_check_response(self, is_last_successful):
    """Sends the health check request and checks the result.

    Args:
      is_last_successful: Whether the last request was successful.

    Returns:
      A bool indicating whether or not the instance is healthy.
    """
    start_response = start_response_utils.CapturingStartResponse()
    try:
      response = self._send_request(start_response, is_last_successful)
    except request_info.Error:
      logging.warning('Health check for instance {instance} is not '
                      'ready yet.'.format(instance=self._instance.instance_id))
      return False
    logging.debug('Health check response %s and status %s for instance %s.',
                  response, start_response.status, self._instance.instance_id)

    return start_response.status.split()[0] == '200'

  def _restart_instance(self):
    """Restarts the running instance, and stops the current health checker."""
    logging.warning('Restarting instance %s because of failed health checks.',
                    self._instance.instance_id)
    self._running = False
    self._restart()
