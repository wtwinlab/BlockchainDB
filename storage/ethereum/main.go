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
	KVStore "github.com/sbip-sg/BlockchainDB/storage/ethereum/contracts/KVStore"
)

func main() {
	// Create a client instance to connect to ganache
	client, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		fmt.Println(err)
	}

	// block := Modules.GetLatestBlock(*client)
	privateKey, err := crypto.HexToECDSA("6845da943907b6cf078d224ab6a48614bae1b3ab554c3a5a7df229a3e4386686")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) //  wei
	auth.GasLimit = uint64(300000)
	auth.GasPrice = big.NewInt(0)

	address := common.HexToAddress("80a391cA469bFdC5d6103890A56DC832401fC0fE")
	instance, err := KVStore.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	key := []byte("tianwen.go")
	value := []byte("hello world")

	tx, err := instance.Set(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	result, err := instance.Get(&bind.TransactOpts{}, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}
