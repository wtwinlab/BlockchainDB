#!/usr/bin/env bash
#set -x

replicaIDs=${1:-1}
shardIDs=${2:-1}
echo "Usage: ./scripts/gen_config.sh 4 1"
echo "Generate config files, replicas: ${replicaIDs}, shards: ${shardIDs}"
dir=$(pwd)
echo $dir
tomlDir="$dir/toml.${shardIDs}.${replicaIDs}"

mkdir -p ${tomlDir}

for (( c=1; c<=${replicaIDs}; c++ ))
do
tomlFile="${tomlDir}/config${c}.toml"
rm -f ${tomlFile}
touch ${tomlFile}
echo "self-id = ${c}" > ${tomlFile}
echo "server-node-addr = \"127.0.0.1:$((50000 + ${c}))\"" >> ${tomlFile}
echo "shard-type = \"ethereum\"" >> ${tomlFile}
echo "shard-number = \"${shardIDs}\"" >> ${tomlFile}
echo "eth-node = \"http://localhost:$((8000 + ${c}))\"" >> ${tomlFile}
echo "eth-hexaddr = \"c6021b15bffcb65c90fc8c52d4ec34e5caa2ae27\"" >> ${tomlFile}
echo "eth-hexkey = \"c60ccf8851c2dc099aace5af7922df16a9cab438d9879dd7c55d0df4f3eb199a\"" >> ${tomlFile}
echo "fab-node = \"127.0.0.1:$((40070 + ${c}))\"" >> ${tomlFile}
echo "fab-config = \"connection${c}.yaml\"" >> ${tomlFile}
echo '' >> ${tomlFile}

echo '# This is the information that each replica is given about the other shards' >> ${tomlFile}
	for (( j=1; j<=${shardIDs}; j++ ))
	do
	echo '[[shards]]' >> ${tomlFile}
	echo "shard-id = ${j}" >> ${tomlFile}
	echo "shard-partition-key = \"eth${j}-\"" >> ${tomlFile}
	echo "shard-type = \"ethereum\"" >> ${tomlFile}
	echo "eth-node = \"http://localhost:$((8000 + ${j}))\"" >> ${tomlFile}
	echo "eth-hexaddr = \"c6021b15bffcb65c90fc8c52d4ec34e5caa2ae27\"" >> ${tomlFile}
	echo "eth-hexkey = \"c60ccf8851c2dc099aace5af7922df16a9cab438d9879dd7c55d0df4f3eb199a\"" >> ${tomlFile}
	echo "fab-node = \"127.0.0.1:$((40070 + ${j}))\"" >> ${tomlFile}
	echo "fab-config = \"connection${j}.yaml\"" >> ${tomlFile}
	echo '' >> ${tomlFile}
	done
echo "Generate config file toml.${replicaIDs}/config${c}.toml"
done


echo "Done!"