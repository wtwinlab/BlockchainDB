#!/usr/bin/env bash
#set -x

shardIDs=${1:-4}
echo "Usage: ./scripts/gen_config.sh 4"
echo "Generate config files, cluster size: ${shardIDs}"
dir=$(pwd)
echo $dir
tomlDir="$dir/toml.${shardIDs}"

mkdir -p ${tomlDir}




for (( c=1; c<=${shardIDs}; c++ ))
do
tomlFile="${tomlDir}/hotstuff${c}.toml"
rm -f ${tomlFile}
touch ${tomlFile}
echo "self-id = ${c}" > ${tomlFile}
echo 'pacemaker = "fixed"' >> ${tomlFile}
# round-robin
echo 'leader-id = 1' >> ${tomlFile}
echo "root-cas = [\"keys.${shardIDs}/ca.crt\"]" >> ${tomlFile}
#echo 'benchmark = true' >> ${tomlFile}
# echo 'tls = true' >> ${tomlFile}
echo 'batch-size = 100' >> ${tomlFile}
#echo "input = \"input.txt\"" >> ${tomlFile}
#echo "print-commands = \"output/result1.txt\"" >> ${tomlFile}
echo "client-listen = \"127.0.0.1:$((20070 + ${c}))\"" >> ${tomlFile}
echo "self-veritas-node = \"127.0.0.1:$((40070 + ${c}))\"" >> ${tomlFile}
echo "self-redis-address = \"127.0.0.1:$((30070 + ${c}))\"" >> ${tomlFile}
echo '' >> ${tomlFile}

echo '# This is the information that each replica is given about the other replicas' >> ${tomlFile}
	for (( j=1; j<=$shardIDs; j++ ))
	do
	echo '[[replicas]]' >> ${tomlFile}
	echo "id = ${j}" >> ${tomlFile}
	echo "peer-address = \"127.0.0.1:$((10070 + ${j}))\"" >> ${tomlFile}
	echo "client-address = \"127.0.0.1:$((20070 + ${j}))\"" >> ${tomlFile}
	echo "redis-address = \"127.0.0.1:$((30070 + ${j}))\"" >> ${tomlFile}
	echo "pubkey = \"keys.${shardIDs}/r${j}.key.pub\"" >> ${tomlFile}
	echo "cert = \"keys.${shardIDs}/r${j}.crt\"" >> ${tomlFile}
	echo '' >> ${tomlFile}
	done
echo "Generate config file toml.${shardIDs}/config${c}.toml"
done


echo "Done!"