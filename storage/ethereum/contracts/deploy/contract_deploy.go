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

	//ethnode := "http://localhost:8000"
	hexkey := "cab9d9e123e4ebe529ad3054e61308fe6411d3eae05b0a16522f0bd92de3c644"

	hexaddress, tx, instance, err := ClientSDK.DeployEthereumKVStoreContract(ethnode, hexkey)
	if err != nil {
		log.Fatal("DeployEthereumKVStoreContract", err)
	}

	fmt.Println("address", hexaddress)
	fmt.Println("tx: ", tx)
	time.Sleep(30 * time.Second)

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
	_ = instance

}
