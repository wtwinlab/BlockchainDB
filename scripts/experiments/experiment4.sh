#!/usr/bin/env bash
#set -x

bestnodes=${1:-16}
bestclients=${2:-16}

# Experiment 4
echo "========================================================"
printf -v date '%(%Y-%m-%d %H:%M:%S)T\n' -1 
echo " Experiment 4 start"
make fast nodes=${bestnodes}
make test nodes=${bestnodes} clients=${bestclients} workload=a
make test nodes=${bestnodes} clients=${bestclients} workload=b
make test nodes=${bestnodes} clients=${bestclients} workload=c
printf -v date '%(%Y-%m-%d %H:%M:%S)T\n' -1 
echo "========================================================"
