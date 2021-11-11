#!/usr/bin/env bash
# set -x

echo "Stop all bcdbnodes"
pgrep -f "bcdbnode" || true
pkill -f "bcdbnode"|| true

echo "##################### Stop bcdbnodes successfully! ##########################"
