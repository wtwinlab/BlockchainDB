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

	ethnode := "http://localhost:8000"
	//hexaddress := "0x062CeC99B1A3e4a5653125cF5ab390C8CE02df7A" //
	//hexaddress := "0x8Cf891EE6F2b232516b6162B1b04D5858B4926a5"
	hexaddress := "0x6005FC466f3e8BC198d192DdDB099fe2998Ba5d1"
	//hexaddress := "0xc6021b15bffcb65c90fc8c52d4ec34e5caa2ae27"
	hexkey := "c60ccf8851c2dc099aace5af7922df16a9cab438d9879dd7c55d0df4f3eb199a"

	ethereumconn, err := ClientSDK.NewEthereumKVStoreInstance(ethnode, hexaddress, hexkey)
	if err != nil {
		log.Fatal(err)
	}

	//key := []byte("tianwen")
	key := "tianwen-1"
	value := "hello world"
	result1, err := ethereumconn.Write(key, value)
	if err != nil {
		log.Fatal("error ethereumconn.Write ", err)
	}
	fmt.Println(result1)
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
