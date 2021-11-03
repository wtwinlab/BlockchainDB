package main

import (
	"fmt"
	"log"

	ClientSDK "github.com/sbip-sg/BlockchainDB/storage/ethereum/clientSDK"
)

func main() {

	ethnode := "http://localhost:8000"
	hexaddress := "b16882db1820b19ea52e7167e68b4ee03b6abb39"
	hexkey := "2243d6f09ad779cc50ccecc3a6779b62ba6c0bb1390e6b0dfcd8fca5eea98193"

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
