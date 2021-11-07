package main

import (
	"fmt"
	"log"

	ClientSDK "github.com/sbip-sg/BlockchainDB/storage/ethereum/clientSDK"
)

func main() {

	ethnode := "http://localhost:8000"
	hexaddress := "c6021b15bffcb65c90fc8c52d4ec34e5caa2ae27"
	hexkey := "c60ccf8851c2dc099aace5af7922df16a9cab438d9879dd7c55d0df4f3eb199a"

	ethereumconn, err := ClientSDK.NewEthereumKVStoreInstance(ethnode, hexaddress, hexkey)
	if err != nil {
		log.Fatal(err)
	}

	key := "tianwen.go"
	// value := "hello world"
	// err = ethereumconn.Write(key, value)
	// if err != nil {
	// 	fmt.Println("error ethereumconn.Write ", err)
	// }

	result, err := ethereumconn.Read(key)
	if err != nil {
		fmt.Println("error ethereumconn.Read ", err)
	}
	fmt.Println(result)

}
