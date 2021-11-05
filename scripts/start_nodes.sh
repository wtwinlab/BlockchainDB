#!/usr/bin/env bash
# set -x

# echo "restart: kill all previous bcdbnode"
# pgrep -f "bcdbnode"
# pkill -f "bcdbnode"
# sleep 5

dir=$(pwd)
bin="$dir/cmd/bcdbnode"

echo "Start blockchaindb server nodes, Please input server node size(default 4)"
replicaIDs=${1:-4}
shardIDs=${2:-1}

for (( c=1; c<=$replicaIDs; c++ ))
do 
$bin --config="toml.${replicaIDs}/config"${c} &
echo "bcdbnode$c start with config file toml.${replicaIDs}.${shardIDs}/config$c.toml"
done

echo "#########################################################################"
echo "##################### Start blockchaindb server nodes successfully! ##########"
echo "#########################################################################"

