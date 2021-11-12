#!/usr/bin/env bash
#set -x


bestnodes=${1:-16}
bestclients=${2:-16}

# Experiment 1
./experiment1.sh 4 > experiment1.log 2>&1


# Experiment 2
./experiment2.sh ${bestclients} > experiment2.log 2>&1


# Experiment 4
./experiment3.sh ${bestnodes} ${bestclients} > experiment4.log 2>&1



# Experiment 3
echo "========================================================"
printf -v date '%(%Y-%m-%d %H:%M:%S)T\n' -1 
echo " Experiment 3 start"
