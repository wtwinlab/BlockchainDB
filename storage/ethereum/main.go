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
	client, err := ethclient.Dial("http://localhost:8000")

	if err != nil {
		fmt.Println(err)
	}

	// block := Modules.GetLatestBlock(*client)
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
	auth.Value = big.NewInt(0) //  wei
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xA2bB64f99C51E51875f920052e1EDF953a83ABd3")
	instance, err := KVStore.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	key := []byte("tianwen.go")
	value := []byte("hello world")
	fmt.Println(key)
	fmt.Println(value)
	tx, err := instance.Set(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	fmt.Println(auth.From.Hex())
	auth.GasPrice = big.NewInt(10)
	result, err := instance.Data(&bind.TransactOpts{From: auth.From,
		Signer: auth.Signer,
		Value:  nil}, key)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)

	result2, err := instance.Get(&bind.TransactOpts{From: auth.From,
		Signer: auth.Signer,
		Value:  nil}, key)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result2)

}
