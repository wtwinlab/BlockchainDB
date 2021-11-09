#!/usr/bin/env bash
# set -x

# echo "restart: kill all previous bcdbnode"
# pgrep -f "bcdbnode"
# pkill -f "bcdbnode"
# sleep 5

dir=$(dirname "$0")

echo "Start blockchaindb server nodes, Please input server node size(default 4)"
replicaIDs=${1:-1}
shardIDs=${2:-1}

bin="$dir/../cmd/bcdbnode/bcdbnode"
tomlDir="$dir/../config.nodes.${shardIDs}.${replicaIDs}"

for (( c=1; c<=$replicaIDs; c++ ))
do 
$bin --config="${tomlDir}/config${c}" &
echo "bcdbnode$c start with config file config.nodes.${replicaIDs}.${shardIDs}/config$c.toml"
done

echo "#########################################################################"
echo "##################### Start blockchaindb server nodes successfully! ##########"
echo "#########################################################################"

