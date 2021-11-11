# BlockchainDB

##### Prepare

```
make download
make build
```



##### 1. Start blockchain network 

(default: ethereum poa)

```
make ethup shards=1 nodes=4

```

##### 2. Start bcdb nodes

```
make install shards=1 nodes=4

```

##### 3. Run ycsb tests

```
make test nodes=4 clients=4
```

Check test result: test.4.4.log
