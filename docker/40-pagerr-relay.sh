#!/bin/sh
# Starts the Quick connect relay in the background before nginx starts.
# If it fails to start, Pagerr still works; Quick connect just hides itself.
set -e
if [ -x /usr/local/bin/pagerr-relay ]; then
  /usr/local/bin/pagerr-relay >/dev/null 2>&1 &
fi
