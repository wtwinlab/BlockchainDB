#!/bin/bash
#args: number_of_nodes, number of networks
#replicaIDs=${1:-1}
set -ex

shardIDs=${1:-1}

cd `dirname ${BASH_SOURCE-$0}`
. env.sh

genesisDir=${ETH_CONFIG}.${shardIDs}
genesisTemplate=${ETH_HOME}/networks/CustomGenesis.template
mkdir -p $genesisDir

echo '# This is custom genesis config template given about each shard'
template=$(<${genesisTemplate})
echo "${template}"

for (( j=1; j<=$shardIDs; j++ ))
do
genesisFile="${genesisDir}/CustomGenesis_${j}.json"
rm -f ${genesisFile}
touch ${genesisFile}
chainIdByShard=$((1000 + ${j}))
jq -n --argjson chainIdByShard $chainIdByShard "${template}" > ${genesisFile}
echo "chainId: $chainIdByShard"
done
echo "Generate genesis file file ${genesisFile}"