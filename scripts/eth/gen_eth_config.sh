#!/bin/bash
#args: number_of_nodes, number of networks
#replicaIDs=${1:-1}
set -ex

shardIDs=${1:-1}
nodeIDs=${2:-4}

cd `dirname ${BASH_SOURCE-$0}`
. env.sh

genesisDir=${ETH_CONFIG}.${shardIDs}.${nodeIDs}
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
signer1=`geth --datadir=${ETH_DATA}_${j}_1 --password <(echo -n "") account new | cut -d '{' -f2 | cut -d '}' -f1`
export BootSignerAddress=${signer1}
extraData="0x0000000000000000000000000000000000000000000000000000000000000000${signer1}0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
cp $genesisTemplate ${genesisFile}
sed -i "s/Signer1/$signer1/" ${genesisFile}
sed -i "s/ChainIdByShard/$chainIdByShard/" ${genesisFile}
sed -i "s/ExtraData/$extraData/" ${genesisFile}

echo "chainId: $chainIdByShard"
done
echo "Generate genesis file file ${genesisFile}"