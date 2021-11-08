package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ClientSDK "github.com/sbip-sg/BlockchainDB/storage/ethereum/clientSDK"
	KVStore "github.com/sbip-sg/BlockchainDB/storage/ethereum/contracts/KVStore"
)

func main() {

	//ethnode := "http://localhost:8000"
	ethnode := "/home/tianwen/Data/eth_1_1/geth.ipc"
	//hexaddress := "0x062CeC99B1A3e4a5653125cF5ab390C8CE02df7A" //
	//hexaddress := "0x8Cf891EE6F2b232516b6162B1b04D5858B4926a5"
	//hexaddress := "01594dbb20c483de0b0a3ed38389fd8a45c70d84"
	hexaddress := "0x8697d21734E20dA53f5fB198F2B6aB1bDa1f2a11"
	hexkey := "cab9d9e123e4ebe529ad3054e61308fe6411d3eae05b0a16522f0bd92de3c644"
	// default account
	//hexaddress := "0xc6021b15bffcb65c90fc8c52d4ec34e5caa2ae27"
	//hexkey := "c60ccf8851c2dc099aace5af7922df16a9cab438d9879dd7c55d0df4f3eb199a"

	ethereumconn, err := ClientSDK.NewEthereumKVStoreInstance(ethnode, hexaddress, hexkey)
	if err != nil {
		log.Fatal(err)
	}

	//key := []byte("tianwen")
	key := "tianwen-2"
	value := "hello world"
	result1, err := ethereumconn.Write(key, value)
	if err != nil {
		log.Fatal("error ethereumconn.Write ", err)
	}
	fmt.Println("write tx: ", result1)
	result, err := ethereumconn.Read(key)
	if err != nil {
		log.Fatal("error ethereumconn.Read ", err)
	}
	fmt.Println(string(result))

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

	contract, err := KVStore.NewStore(address, client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n", err)
	}
	privateKey, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		log.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice       // big.NewInt(0)
	result2, err := contract.Get(auth, []byte(key))
	if err != nil {
		log.Fatal("error ethereumconn.Read ", err)
	}
	fmt.Println(string(result2.Data()))

}
