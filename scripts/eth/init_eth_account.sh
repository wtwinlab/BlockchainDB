#!/bin/bash
#
#args: number of networks, number_of_nodes
#
set -ex

shardIDs=${1:-1}
nodeIDs=${2:-1}

cd `dirname ${BASH_SOURCE-$0}`
. env.sh
genesisDir=${ETH_CONFIG}.${shardIDs}

for (( j=1; j<=${shardIDs}; j++ ))
do
genesisFile="${genesisDir}/CustomGenesis_${j}.json"
  for (( i=1; i<=${nodeIDs}; i++ ))
  do
  echo "Using custom genesis file: ${genesisFile}, datadir: ${ETH_DATA}_${j}_${i}"
  geth --datadir=${ETH_DATA}_${j}_${i} init ${genesisFile}
  geth --datadir=${ETH_DATA}_${j}_${i} --password <(echo -n "") account new
  #geth --rpc --rpcport "8085" --datadir Data/TestChain init genesis.json
  done
done