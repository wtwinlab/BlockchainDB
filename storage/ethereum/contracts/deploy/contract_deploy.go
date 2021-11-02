package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	KVStore "github.com/sbip-sg/BlockchainDB/storage/ethereum/contracts/KVStore"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("57ad40bdff72011930654898a549424cbb6d78ecefbfef64335c2dbc3453e097")
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

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice       //big.NewInt(0)
	fmt.Println(gasPrice)
	address, tx, instance, err := KVStore.DeployStore(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())

	// test contract
	store, err := KVStore.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	key := []byte("tianwen.go")
	value := []byte("hello world")
	nonce2, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce2))
	gasPrice2, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice2
	fmt.Println(gasPrice2)
	tx, err = store.Set(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	result, err := store.Get(&bind.TransactOpts{}, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	_ = instance
}
