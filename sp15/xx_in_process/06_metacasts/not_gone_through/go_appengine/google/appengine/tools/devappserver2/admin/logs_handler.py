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
"""A handler that displays logs for instances."""

import json
import os

import google
import requests

from google.appengine.tools.devappserver2 import log_manager
from google.appengine.tools.devappserver2.admin import admin_request_handler


class LogsHandler(admin_request_handler.AdminRequestHandler):
  _REQUIRED_PARAMS = ['app', 'module', 'version', 'instance', 'log_type']

  def get(self):
    try:
      ps = self.request.params
      params = dict([(p, ps[p]) for p in self._REQUIRED_PARAMS])
    except KeyError, e:
      self.abort(404, detail='Missing log request parameter %s.' % e.message)

    # Forward request to LogServer
    host = os.environ.get(log_manager.APP_ENGINE_LOG_SERVER_HOST)
    port = os.environ.get(log_manager.APP_ENGINE_LOG_SERVER_PORT)
    if not host or not port:
      self.abort(404, detail='LogServer Host and Port must be set')

    r = requests.get('http://{host}:{port}'.format(host=host, port=port),
                     params=params)
    params['logs'] = json.loads(r.text)

    self.response.write(self.render('instance_logs.html', params))
