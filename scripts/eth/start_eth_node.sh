#!/bin/bash
shardID=${1:-1}

cd `dirname ${BASH_SOURCE-$0}`
. env.sh
#geth --datadir=$ETH_DATA --rpc --rpcaddr 0.0.0.0 --rpcport "8000" --rpccorsdomain "*" --gasprice 0 --networkid 9119 --unlock 0 --password <(echo -n "") js <(echo 'console.log(admin.nodeInfo.enode);') 2>/dev/null |grep enode | perl -pe "s/\[\:\:\]/$ip_addr/g" | perl -pe "s/^/\"/; s/\s*$/\"/;"

#geth --datadir=$ETH_DATA --rpc --rpcaddr 0.0.0.0 --rpcport "8000" --rpccorsdomain "*" --gasprice 0 --networkid 9119 --unlock 0 --password <(echo -n "") js <(echo 'console.log(admin.nodeInfo.enode);') 

${ETH_BIN}/geth --datadir=${ETH_DATA}_${shardID} --rpc --rpcport "8000" --syncmode "full" --cache 4096 --gasprice 0 --networkid 10001 --nodiscover --unlock 0 #--password <(echo -n "") js <(echo 'console.log(admin.nodeInfo.enode);') 
#--mine --minerthreads 1 --unlock 0 console 2>>geth.log
# --cache 4096
#${ETH_BIN}/geth --datadir=${ETH_DATA}_${shardID} --rpc --rpcport \"$((8000 + ${c}))\" --syncmode "full" --gasprice 0 --networkid 10001 --nodiscover --unlock 0 #--password <(echo -n "") js <(echo 'console.log(admin.nodeInfo.enode);') 
