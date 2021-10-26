package connectors

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

type EthereumConnector struct {
	client *ethclient.Client
	kv     *KVStore.Store
	auth   *bind.TransactOpts
}

func NewKVStoreInstance(addressHex string, ethnode string) (*EthereumConnector, error) {
	var ethereumConn *EthereumConnector
	address := common.HexToAddress(addressHex)
	client, err := ethclient.Dial(ethnode)
	if err != nil {
		fmt.Println(err)
	}
	ethereumConn.client = client
	store, err := KVStore.NewStore(address, client)
	if err != nil {
		return nil, err
	}
	ethereumConn.kv = store
	return ethereumConn, nil

}

func (ethereumConn *EthereumConnector) Get(key []byte) (string, error) {
	//value, err := ethereumConn.kv.Get(&bind.TransactOpts{}, key)
	return "", nil
}

func (ethereumConn *EthereumConnector) Set(key, value []byte) error {
	gasPrice, err := ethereumConn.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("65f52ec60aadd071de2d5a9241f5d80d6edd159dbd0aebaeec9699e76abe6653")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethereumConn.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	tx, err := ethereumConn.kv.Set(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent
	return nil
}

func (ethereum *EthereumConnector) Verify(key, value string) error {

	return nil
}
