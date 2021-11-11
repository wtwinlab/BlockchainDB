package main

import (
	"context"
	"fmt"

	ClientSDK "github.com/sbip-sg/BlockchainDB/storage/fabric/clientSDK"
)

func main() {

	fabconn, err := ClientSDK.NewFabricKVStoreInstance()
	if err != nil {
		fmt.Println("Failed to NewFabricKVStoreInstance", err)
	}

	key := "tianwen"
	value := "hello world"

	tx, err := fabconn.Write(context.Background(), key, value)
	if err != nil {
		fmt.Printf("Failed to Write: %v", err)
	}
	fmt.Println(tx)

	val, err := fabconn.Read(context.Background(), key)
	if err != nil {
		fmt.Printf("Failed to Read: %v", err)
	}
	fmt.Println(val)
}
