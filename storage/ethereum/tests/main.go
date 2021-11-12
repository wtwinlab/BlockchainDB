package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sbip-sg/BlockchainDB/bcdbnode/config"
	ClientSDK "github.com/sbip-sg/BlockchainDB/storage/ethereum/clientSDK"
)

func main() {
	//ganache
	// ethnode := "http://localhost:7545"
	// hexaddress := "0xa5385Dbde6EA7bEa767213db23B9A6D39b912a33" //contract address
	// hexkey := "98670003f6f38d426a6b76cd7143926b9af17391a5ff76f8ee14f162511d6814"

	//local eth_1_1
	// ethnode := "/home/tianwen/Data/eth_1_1/geth.ipc"
	// hexaddress := "0x0803521274Fb66b54Ef6CF22A801713B1299b5cD" //"0x8c60fa428050904a19da95a07e8eb3e382bd2329"
	// hexkey := "5c01c3481aadd0a1d0828c61e32b2af82724b844de1b06673fa246df61d1a887"
	//redisAddr := "127.0.0.1:60001"

	//config from file
	configFile := flag.String("config", "config.nodes.1.4/config_1_1", "The path to the config file")
	flag.Parse()
	var conf config.Options
	err := config.ReadConfig(&conf, *configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read config: %v\n", err)
		os.Exit(1)
	}
	ethnode := conf.Shards[0].EthNode
	hexkey := conf.Shards[0].EthHexKey
	hexaddress := conf.Shards[0].EthHexAddr
	redisAddr := conf.Shards[0].RedisAddr

	ethereumconn, err := ClientSDK.NewEthereumKVStoreInstance(ethnode, hexaddress, hexkey, redisAddr)
	if err != nil {
		log.Fatal(err)
	}

	key := "tianwen-7"
	value := "helloworld7"
	result1, err := ethereumconn.Write(context.Background(), key, value)
	if err != nil {
		log.Fatal("error ethereumconn.Write ", err)
	}
	fmt.Println("write tx: ", result1)
	time.Sleep(5 * time.Second)

	result, err := ethereumconn.Read(context.Background(), key)
	if err != nil {
		log.Fatal("error ethereumconn.Read ", err)
	}
	fmt.Println(string(result))

	result2, err := ethereumconn.Verify(context.Background(), "set", key)
	fmt.Println(result2)

	os.Exit(0)
	//Debug
	client, err := ethclient.Dial(ethnode)

	if err != nil {
		fmt.Println("error ethclient Dail "+ethnode, err)
	}
	address := common.HexToAddress(hexaddress)
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract) // is contract: true

}
