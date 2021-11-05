#!/usr/bin/env bash
#set -x

replicaIDs=${1:-4}
shardIDs=${2:-1}
echo "Usage: ./scripts/gen_config.sh 4 1"
echo "Generate config files, cluster size: ${shardIDs}"
dir=$(pwd)
echo $dir
tomlDir="$dir/toml.${shardIDs}"

mkdir -p ${tomlDir}




for (( c=1; c<=${replicaIDs}; c++ ))
do
tomlFile="${tomlDir}/config${c}.toml"
rm -f ${tomlFile}
touch ${tomlFile}
echo "self-id = ${c}" > ${tomlFile}
echo 'server-node-addr = 1' >> ${tomlFile}
echo "shard-type = \"eth\"" >> ${tomlFile}
echo "eth-node = \"http://localhost:$((8000 + ${c}))\"" >> ${tomlFile}
echo "eth-hexaddr = \"b16882db1820b19ea52e7167e68b4ee03b6abb39\"" >> ${tomlFile}
echo "eth-hexkey = \"2243d6f09ad779cc50ccecc3a6779b62ba6c0bb1390e6b0dfcd8fca5eea98193\"" >> ${tomlFile}
echo "fab-node = \"127.0.0.1:$((40070 + ${c}))\"" >> ${tomlFile}
echo "fab-config = \"connection${c}.yaml\"" >> ${tomlFile}
echo '' >> ${tomlFile}

echo '# This is the information that each replica is given about the other shards' >> ${tomlFile}
	for (( j=1; j<=$shardIDs; j++ ))
	do
	echo '[[shards]]' >> ${tomlFile}
	echo "shard-id = ${j}" >> ${tomlFile}
	echo "shard-partition-key = \"eth${j}-\"" >> ${tomlFile}
	echo "shard-type = \"eth\"" >> ${tomlFile}
	echo "eth-node = \"http://localhost:$((8000 + ${j}))\"" >> ${tomlFile}
	echo "eth-hexaddr = \"b16882db1820b19ea52e7167e68b4ee03b6abb39\"" >> ${tomlFile}
	echo "eth-hexkey = \"2243d6f09ad779cc50ccecc3a6779b62ba6c0bb1390e6b0dfcd8fca5eea98193\"" >> ${tomlFile}
	echo "fab-node = \"127.0.0.1:$((40070 + ${j}))\"" >> ${tomlFile}
	echo "fab-config = \"connection${j}.yaml\"" >> ${tomlFile}
	echo '' >> ${tomlFile}
	done
echo "Generate config file toml.${replicaIDs}/config${c}.toml"
done


echo "Done!"