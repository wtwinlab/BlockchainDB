#!/bin/bash
#

#./network.sh -h
./network.sh down
./network.sh up createChannel
docker ps -a
./network.sh deployCC -ccn basic -ccp ../asset-transfer-basic/chaincode-go -ccl go

./network.sh createChannel -c kvchannel
./network.sh deployCC -c kvchannel -ccn kvstore -ccp ../kvstore/chaincode-go -ccl go

echo "=================== Success ==================="
