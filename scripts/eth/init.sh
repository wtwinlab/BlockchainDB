#!/bin/bash
#args: number of networks, number_of_nodes
set -ex

shardIDs=${1:-1}
genesisDir=${ETH_CONFIG}.${shardIDs}

cd `dirname ${BASH_SOURCE-$0}`
. env.sh

for (( j=1; j<=$shardIDs; j++ ))
do
genesisFile="${genesisDir}/CustomGenesis_${j}.json"
echo "Using custom genesis file: ${genesisFile}"
geth --datadir=$ETH_DATA init ${genesisFile}
geth --datadir=$ETH_DATA --password <(echo -n "") account new
#geth --rpc --rpcport "8085" --datadir Data/TestChain init genesis.json

done