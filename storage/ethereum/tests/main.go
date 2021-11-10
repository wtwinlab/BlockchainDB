package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ClientSDK "github.com/sbip-sg/BlockchainDB/storage/ethereum/clientSDK"
)

func main() {
	//ganache
	ethnode := "http://localhost:7545"
	hexaddress := "0xa5385Dbde6EA7bEa767213db23B9A6D39b912a33" //contract address
	hexkey := "98670003f6f38d426a6b76cd7143926b9af17391a5ff76f8ee14f162511d6814"

	//local eth_1_1
	// ethnode := "/home/tianwen/Data/eth_1_1/geth.ipc"
	// hexaddress := "0x70fa2c27a4e365cdf64b2d8a6c36121eb80bb442"
	// hexkey := "35fc8e4f2065b6813078a08069e3a946f203029ce2bc6a62339d30c37f978403"

	// default account
	//hexaddress := "0xc6021b15bffcb65c90fc8c52d4ec34e5caa2ae27"
	//hexkey := "c60ccf8851c2dc099aace5af7922df16a9cab438d9879dd7c55d0df4f3eb199a"
	redisAddr := "127.0.0.1:60001"
	ethereumconn, err := ClientSDK.NewEthereumKVStoreInstance(ethnode, hexaddress, hexkey, redisAddr)
	if err != nil {
		log.Fatal(err)
	}

	//key := []byte("tianwen")
	key := "tianwen-6"
	value := "helloworld"
	result1, err := ethereumconn.Write(context.Background(), key, value)
	if err != nil {
		log.Fatal("error ethereumconn.Write ", err)
	}
	fmt.Println("write tx: ", result1)
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
