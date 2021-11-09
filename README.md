# BlockchainDB

##### Start blockchain network

```
//Install Geth@v1.8.23
//reset if exist  ~/Data/eth/
./scripts/eth/gen_eth_config.sh 
./scripts/eth/init_eth_account.sh
./scripts/eth/start_eth_node.sh

//Deploy KVStore contract (storage/ethereum/contracts/KVStore)
go run storage/ethereum/contracts/deploy/contract_deploy.go
```

##### Start bcdb nodes

```
./scripts/gen_config.sh ${shardIDs} ${replicaIDs}
./scripts/gen_ycsb_data.sh

make clean
make build
make install $(nodes)
make test $(nodes) $(clients)
```

