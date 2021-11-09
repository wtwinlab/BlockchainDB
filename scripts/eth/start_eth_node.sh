#!/bin/bash
set -ex

shardID=${1:-1}
nodeID=${2:-1}

cd `dirname ${BASH_SOURCE-$0}`
. env.sh

#${ETH_BIN}/geth --datadir=${ETH_DATA}_${nodeID} --rpc --rpcport "8000" --syncmode "full" --cache 4096 --gasprice 0 --networkid 10001 --mine --minerthreads 1 --unlock 0 console 2> ${ETH_DATA}_${nodeID}/geth.log
#--password <(echo -n "") js <(echo 'console.log(admin.nodeInfo.enode);') 
#--nodiscover 
#--targetgaslimit '67219750000000'

# console
# geth attach ./geth.ipc 

${ETH_BIN}/geth --datadir=${ETH_DATA}_${shardID}_${nodeID}  --rpc --rpcport "$((9000 + ${nodeID} + 1000*${shardID}))" --port "$((30303 + ${nodeID} + 1000*(${shardID}-1)))" --syncmode "full" --cache 4096 --gasprice 0 --networkid $((1000 + ${shardID})) --mine --minerthreads 1 --unlock 0 --password <(echo -n "") 2> ${ETH_DATA}_${shardID}_${nodeID}/geth.log &
