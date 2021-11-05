package clientSDK

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

	BlockchainConnector "github.com/sbip-sg/BlockchainDB/blockchainconnectors/ethereumconnector"
)

func NewEthereumKVStoreInstance(ethnode string, hexaddress string, hexkey string) (*BlockchainConnector.EthereumConnector, error) {

	var conn *BlockchainConnector.EthereumConnector
	client, err := ethclient.Dial(ethnode)

	if err != nil {
		fmt.Println("error ethclient Dail "+ethnode, err)
		return conn, err
	}

	address := common.HexToAddress(hexaddress)
	instance, err := KVStore.NewStore(address, client)
	if err != nil {
		log.Fatal("error create a new instance of Store, bound to a specific deployed contract.", err)
		return conn, err
	}

	conn = &BlockchainConnector.EthereumConnector{Client: client, KV: instance, Hexkey: hexkey}
	return conn, nil
}

func DeployEthereumKVStoreContract(ethnode string, hexkey string) (string, string, *KVStore.Store, error) {
	client, err := ethclient.Dial(ethnode)
	if err != nil {
		log.Println(err)
		return "", "", nil, err
	}

	privateKey, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		log.Println(err)
		return "", "", nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("error casting public key to ECDSA")
		return "", "", nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return "", "", nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return "", "", nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice       //big.NewInt(0)
	fmt.Println("gasPrice ", gasPrice)
	address, tx, instance, err := KVStore.DeployStore(auth, client)
	if err != nil {
		log.Println("KVStore.DeployStore", err)
		return "", "", nil, err
	}

	return address.Hex(), tx.Hash().Hex(), instance, nil
}