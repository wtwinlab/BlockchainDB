#!/usr/bin/env bash
# set -x

echo "Stop all bcdbnodes"
# pgrep -f "bcdbnode" || true
# pkill -f "bcdbnode"|| true
kill -9 $(ps -ef|grep "bcdbnode"|grep -v "grep"|awk '{print $2}')
sleep 4
echo "##################### Stop bcdbnodes successfully! ##########################"
