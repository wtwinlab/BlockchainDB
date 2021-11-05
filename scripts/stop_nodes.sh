#!/usr/bin/env bash
# set -x

echo "Stop all bcdbnodes"
pgrep -f "bcdbnode"
pkill -f "bcdbnode"

echo "##################### Stop nodes successfully! ##########################"
