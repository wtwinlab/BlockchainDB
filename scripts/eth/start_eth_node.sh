#!/bin/bash
#ip_addr=$1
. env.sh
#/home/tianwen/go/src/github.com/ethereum/go-ethereum/build/bin/geth --datadir=$ETH_DATA --rpc --rpcaddr 0.0.0.0 --rpcport "8000" --rpccorsdomain "*" --gasprice 0 --networkid 9119 --unlock 0 --password <(echo -n "") js <(echo 'console.log(admin.nodeInfo.enode);') 2>/dev/null |grep enode | perl -pe "s/\[\:\:\]/$ip_addr/g" | perl -pe "s/^/\"/; s/\s*$/\"/;"

#/home/tianwen/go/src/github.com/ethereum/go-ethereum/build/bin/geth --datadir=$ETH_DATA --rpc --rpcaddr 0.0.0.0 --rpcport "8000" --rpccorsdomain "*" --gasprice 0 --networkid 9119 --unlock 0 --password <(echo -n "") js <(echo 'console.log(admin.nodeInfo.enode);') 

/home/tianwen/go/src/github.com/ethereum/go-ethereum/build/bin/geth --datadir=$ETH_DATA --rpc --rpcport "8000" --gasprice 0 --networkid 9119 --nodiscover #--unlock 0 --password <(echo -n "") js <(echo 'console.log(admin.nodeInfo.enode);') 