package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ClientSDK "github.com/sbip-sg/BlockchainDB/storage/ethereum/clientSDK"
)

func main() {

	//local
	ethnode := "/home/tianwen/Data/eth_1_1/geth.ipc"
	hexkey := "35fc8e4f2065b6813078a08069e3a946f203029ce2bc6a62339d30c37f978403"

	// ethnode := "http://localhost:7545"
	// hexkey := "2b2f78b59cb38b1de758dcdbe71cea29d9d4907e6921a3dd77a8e56fc71de8bf"

	hexaddress, tx, _, err := ClientSDK.DeployEthereumKVStoreContract(ethnode, hexkey)
	if err != nil {
		log.Fatal("DeployEthereumKVStoreContract", err)
	}

	fmt.Println("address", hexaddress)
	fmt.Println("tx: ", tx)
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
	fmt.Printf("is contract: %v\n", isContract) // is contract: true

	// key := []byte("tianwen")
	// value := []byte("hello world")
	// result1, err := instance.Set(key, value)
	// if err != nil {
	// 	log.Fatal("error ethereumconn.Write ", err)
	// }
	// fmt.Println(result1)
	// key := []byte("tianwen")
	// result, err := instance.Get(nil, key)
	// if err != nil {
	// 	log.Fatal("error instance.Get ", err)
	// }
	// fmt.Println(string(result.Data()))
	//_ = instance

}
