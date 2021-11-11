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

	//local
	// ethnode := "/home/tianwen/Data/eth_1_1/geth.ipc"
	// hexkey := "5c01c3481aadd0a1d0828c61e32b2af82724b844de1b06673fa246df61d1a887"

	//ganache
	// ethnode := "http://localhost:7545"
	// hexkey := "98670003f6f38d426a6b76cd7143926b9af17391a5ff76f8ee14f162511d6814"

	//config from file
	configFile := flag.String("config", "config.eth.1.4/shard_1", "The path to the config file")
	flag.Parse()
	var conf config.Options
	err := config.ReadConfig(&conf, *configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read config: %v\n", err)
		os.Exit(1)
	}
	ethnode := conf.EthNode
	hexkey := conf.EthHexKey

	hexaddress, tx, _, err := ClientSDK.DeployEthereumKVStoreContract(ethnode, hexkey)
	if err != nil {
		log.Fatal("DeployEthereumKVStoreContract", err)
	}

	fmt.Printf("eth-hexaddr = \"%v\"\n", hexaddress)
	fmt.Printf("contract-tx = \"%v\"\n", tx)
	time.Sleep(10 * time.Second)

	//Debug
	client, err := ethclient.Dial(ethnode)
	if err != nil {
		fmt.Println("error ethclient Dail "+ethnode, err)
	}
	//hexaddress := "0xdC5F73E49DdC7ECF506bEAEB4f53670bC6ac53c6"
	address := common.HexToAddress(hexaddress)
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0
	fmt.Printf("contract = %v\n", isContract) // is contract: true
}
