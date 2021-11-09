#!/usr/bin/env bash
#set -x



shardIDs=${1:-1}

#docker pull redis:latest

docker rm -f $(sudo -S docker ps -aq  --filter ancestor=redis)

for (( c=1; c<=${shardIDs}; c++ ))
do
docker run -itd --name shard${c}-redis -p $((60000 + ${c})):6379 redis
echo "redis db start with port $((60000 + ${c}))"
done

echo "##################### Start redis dbs successfully! #####################"