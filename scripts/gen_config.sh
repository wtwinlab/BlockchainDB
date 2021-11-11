#!/usr/bin/env bash
set -ex

shardIDs=${1:-1}
replicaIDs=${2:-1}
. env.sh
echo "Usage: ./scripts/gen_config.sh 4 1"
echo "Generate config files, shards: ${shardIDs}, replicas: ${replicaIDs}"
dir=$(pwd)
echo $dir
tomlDir="$dir/config.nodes.${shardIDs}.${replicaIDs}"

mkdir -p ${tomlDir}
rm -rf ETH_DATA=$HOME/Data/eth*

for (( c=1; c<=${replicaIDs}; c++ ))
do
tomlFile="${tomlDir}/config${c}.toml"
rm -f ${tomlFile}
touch ${tomlFile}
echo "self-id = ${c}" > ${tomlFile}
echo "server-node-addr = \"127.0.0.1:$((50000 + ${c}))\"" >> ${tomlFile}
echo "shard-type = \"ethereum\"" >> ${tomlFile}
echo "shard-number = \"${shardIDs}\"" >> ${tomlFile}
echo "eth-node = \"~/Data/eth_${shardIDs}_${c}/geth.ipc\"" >> ${tomlFile}
echo "eth-hexaddr = \"0x70fa2c27a4e365cdf64b2d8a6c36121eb80bb442\"" >> ${tomlFile}
echo "eth-hexkey = \"35fc8e4f2065b6813078a08069e3a946f203029ce2bc6a62339d30c37f978403\"" >> ${tomlFile}
echo "fab-node = \"127.0.0.1:$((40000 + ${c}))\"" >> ${tomlFile}
echo "fab-config = \"connection${c}.yaml\"" >> ${tomlFile}
echo '' >> ${tomlFile}

echo '# This is the information that each replica is given about the other shards' >> ${tomlFile}
	for (( j=1; j<=${shardIDs}; j++ ))
	do
	echo '[[shards]]' >> ${tomlFile}
	echo "shard-id = ${j}" >> ${tomlFile}
	echo "shard-partition-key = \"eth${j}-\"" >> ${tomlFile}
	echo "shard-type = \"ethereum\"" >> ${tomlFile}
	#echo "eth-node = \"http://localhost:$((9000 + ${c} + 1000*${j}))\"" >> ${tomlFile}
	echo "eth-node = \"$HOME/Data/eth_${shardIDs}_${c}/geth.ipc\"" >> ${tomlFile}
	echo "eth-hexaddr = \"0x70fa2c27a4e365cdf64b2d8a6c36121eb80bb442\"" >> ${tomlFile}
	echo "eth-hexkey = \"35fc8e4f2065b6813078a08069e3a946f203029ce2bc6a62339d30c37f978403\"" >> ${tomlFile}
	echo "fab-node = \"127.0.0.1:$((40000 + ${j}))\"" >> ${tomlFile}
	echo "fab-config = \"connection${j}.yaml\"" >> ${tomlFile}
	echo "redis-address = \"127.0.0.1:$((60000 + ${j}))\"" >> ${tomlFile}
	echo '' >> ${tomlFile}
	done
echo "Generate config file ${tomlFile}"
done


echo "Done!"