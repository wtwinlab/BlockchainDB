package connectors

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

type EthereumConnector struct {
	Client *ethclient.Client
	KV     *KVStore.Store
	Auth   *bind.TransactOpts
	Hexkey string
}

func (ethereumConn *EthereumConnector) Read(key string) (string, error) {
	auth, err := ethereumConn.bindTransactOpts()
	result, err := ethereumConn.KV.Get(auth, []byte(key))
	if err != nil {
		log.Println("error Read", err)
		return "", err
	}
	//fmt.Println(string(result.Data()))
	return string(result.Data()), nil
}

func (ethereumConn *EthereumConnector) Write(key, value string) error {

	auth, err := ethereumConn.bindTransactOpts()
	tx, err := ethereumConn.KV.Set(auth, []byte(key), []byte(value))
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent
	return nil
}

func (ethereumConn *EthereumConnector) bindTransactOpts() (*bind.TransactOpts, error) {
	gasPrice, err := ethereumConn.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println("error parse a secp256k1 private key.", err)
		return nil, err
	}
	privateKey, err := crypto.HexToECDSA(ethereumConn.Hexkey)
	if err != nil {
		log.Println("error casting public key to ECDSA.", err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethereumConn.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println("error return the account nonce of the given account in the pending state.", err)
		return nil, err
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

// func (ethereumConn *EthereumConnector) Get(key []byte) (string, error) {
// 	//value, err := ethereumConn.kv.Get(&bind.TransactOpts{}, key)
// 	return "", nil
// }

// func (ethereumConn *EthereumConnector) Set(key, value []byte) error {
// 	gasPrice, err := ethereumConn.Client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	privateKey, err := crypto.HexToECDSA("65f52ec60aadd071de2d5a9241f5d80d6edd159dbd0aebaeec9699e76abe6653")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("error casting public key to ECDSA")
// 	}

// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	nonce, err := ethereumConn.Client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	auth := bind.NewKeyedTransactor(privateKey)
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)     // in wei
// 	auth.GasLimit = uint64(300000) // in units
// 	auth.GasPrice = gasPrice
// 	tx, err := ethereumConn.KV.Set(auth, key, value)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent
// 	return nil
// }
//
// func NewKVStoreInstance(addressHex string, ethnode string) (*EthereumConnector, error) {
// 	var ethereumConn *EthereumConnector
// 	address := common.HexToAddress(addressHex)
// 	client, err := ethclient.Dial(ethnode)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	ethereumConn.Client = client
// 	store, err := KVStore.NewStore(address, client)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ethereumConn.KV = store
// 	return ethereumConn, nil

// }
