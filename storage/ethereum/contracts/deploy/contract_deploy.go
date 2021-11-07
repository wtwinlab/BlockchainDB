package main

import (
	"fmt"
	"log"

	ClientSDK "github.com/sbip-sg/BlockchainDB/storage/ethereum/clientSDK"
)

func main() {

	ethnode := "http://localhost:8000"
	hexkey := "c60ccf8851c2dc099aace5af7922df16a9cab438d9879dd7c55d0df4f3eb199a"

	address, tx, instance, err := ClientSDK.DeployEthereumKVStoreContract(ethnode, hexkey)
	if err != nil {
		log.Fatal("DeployEthereumKVStoreContract", err)
	}

	fmt.Println("address", address)
	fmt.Println("tx: ", tx)
	//key := "tianwen"
	// value := "hello world"
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
