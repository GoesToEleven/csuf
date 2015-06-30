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
"""Directly processes text of cron.xml.

CronXmlParser is called with an XML string to produce a CronXml object
containing the data from the XML.

CronXmlParser: converts XML to CronXml objct
Cron: describes a single cron specified in cron.xml
"""

from xml.etree import ElementTree

from google.appengine.cron import groc
from google.appengine.cron import groctimespecification
from google.appengine.tools import xml_parser_utils
from google.appengine.tools.app_engine_config_exception import AppEngineConfigException


def GetCronYaml(unused_application, cron_xml_str):
  return _MakeCronListIntoYaml(CronXmlParser().ProcessXml(cron_xml_str))


def _MakeCronListIntoYaml(cron_list):
  """Converts list of yaml statements describing cron jobs into a string."""
  statements = ['cron:']
  for cron in cron_list:
    statements += cron.ToYaml()
  return '\n'.join(statements) + '\n'


class CronXmlParser(object):
  """Provides logic for walking down XML tree and pulling data."""

  def ProcessXml(self, xml_str):
    """Parses XML string and returns object representation of relevant info.

    Args:
      xml_str: The XML string.
    Returns:
      A list of Cron objects containing information about cron jobs from the
      XML.
    Raises:
      AppEngineConfigException: In case of malformed XML or illegal inputs.
    """

    try:
      self.crons = []
      self.errors = []
      xml_root = ElementTree.fromstring(xml_str)
      if xml_root.tag != 'cronentries':
        raise AppEngineConfigException('Root tag must be <cronentries>')

      for child in xml_root.getchildren():
        self.ProcessCronNode(child)

      if self.errors:
        raise AppEngineConfigException('\n'.join(self.errors))

      return self.crons
    except ElementTree.ParseError:
      raise AppEngineConfigException('Bad input -- not valid XML')

  def ProcessCronNode(self, node):
    """Processes XML <cron> nodes into Cron objects.

    The following information is parsed out:
      description: Describing the purpose of the cron job.
      url: The location of the script.
      schedule: Written in groc; the schedule according to which the job is
        executed.
      timezone: The timezone that the schedule runs in.
      target: Which version of the app this applies to.

    Args:
      node: <cron> XML node in cron.xml.
    """
    tag = xml_parser_utils.GetTag(node)
    if tag != 'cron':
      self.errors.append('Unrecognized node: <%s>' % tag)
      return

    cron = Cron()
    cron.url = xml_parser_utils.GetChildNodeText(node, 'url')
    cron.timezone = xml_parser_utils.GetChildNodeText(node, 'timezone')
    cron.target = xml_parser_utils.GetChildNodeText(node, 'target')
    cron.description = xml_parser_utils.GetChildNodeText(node, 'description')
    cron.schedule = xml_parser_utils.GetChildNodeText(node, 'schedule')

    validation_error = self._ValidateCronEntry(cron)
    if validation_error:
      self.errors.append(validation_error)
    else:
      self.crons.append(cron)

  def _ValidateCronEntry(self, cron):


    if not cron.url:
      return 'No URL for <cron> entry'
    if not cron.schedule:
      return "No schedule provided for <cron> entry with URL '%s'" % cron.url
    try:
      groctimespecification.GrocTimeSpecification(cron.schedule)
    except groc.GrocException:
      return ("Text '%s' in <schedule> node failed to parse,"
              ' for <cron> entry with url %s.'
              % (cron.schedule, cron.url))


class Cron(object):
  """Instances contain information about individual cron entries."""
  TZ_GMT = 'UTC'

  def ToYaml(self):
    """Returns data from Cron object as a list of Yaml statements."""
    statements = [
        '- url: %s' % self._SanitizeForYaml(self.url),
        '  schedule: %s' % self._SanitizeForYaml(self.schedule)]
    for optional in ('target', 'timezone', 'description'):
      field = getattr(self, optional)
      if field:
        statements.append('  %s: %s' % (optional, self._SanitizeForYaml(field)))
    return statements

  def _SanitizeForYaml(self, field):
    return "'%s'" % field.replace('\n', ' ').replace("'", "''")
