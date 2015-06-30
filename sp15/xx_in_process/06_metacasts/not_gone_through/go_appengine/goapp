#!/usr/bin/env python
#
# Copyright 2011 Google Inc. All rights reserved.
# Use of this source code is governed by the Apache 2.0
# license that can be found in the LICENSE file.

"""Convenience wrapper for starting a Go tool in the App Engine SDK."""

import argparse
import os
import sys

SDK_BASE = os.path.abspath(os.path.dirname(os.path.realpath(__file__)))
GOROOT = os.path.join(SDK_BASE, 'goroot')


def GetArgsAndArgv():
  """Parses command-line arguments and strips out flags the control this file.

  Returns:
    Tuple of (flags, rem) where "flags" is a map of key,value flags pairs
    and "rem" is a list that strips the used flags and acts as a new
    replacement for sys.argv.
  """
  parser = argparse.ArgumentParser(add_help=False)
  parser.add_argument(
      '--dev_appserver', default=os.path.join(SDK_BASE, 'dev_appserver.py'))
  namespace, rest = parser.parse_known_args()
  flags = vars(namespace)
  rem = [sys.argv[0]]
  rem.extend(rest)
  return flags, rem


if __name__ == '__main__':
  vals, new_argv = GetArgsAndArgv()
  tool = os.path.basename(__file__)
  bin = os.path.join(GOROOT, 'bin', tool)
  os.environ['GOROOT'] = GOROOT
  os.environ['APPENGINE_DEV_APPSERVER'] = vals['dev_appserver']

  # Remove env variables that may be incompatible with the SDK.
  for e in ('GOARCH', 'GOBIN', 'GOOS'):
    os.environ.pop(e, None)

  # Set a GOPATH if one is not set.
  if not os.environ.get('GOPATH'):
    os.environ['GOPATH'] = os.path.join(SDK_BASE, 'gopath')

  os.execve(bin, new_argv, os.environ)
