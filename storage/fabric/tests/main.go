/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
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
	err = fabconn.Write(key, value)
	if err != nil {
		fmt.Printf("Failed to Write: %v", err)
	}

	val, err := fabconn.Read(key)
	if err != nil {
		fmt.Printf("Failed to Read: %v", err)
	}
	fmt.Println(val)
}
