/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	BlockchainConnector "github.com/sbip-sg/BlockchainDB/connectors"
)

func main() {
	var fabconn *BlockchainConnector.FabricConnector
	fabconn, err := NewFabricKVStoreInstance()
	if err != nil {
		fmt.Println(err)
	}
	key := "tianwen"
	value := "world"
	res, err := fabconn.Write(key, value)
	if err != nil {
		fmt.Printf("Failed to Write: %v", err)
	}
	fmt.Println(res)
	val, err := fabconn.Read(key)
	if err != nil {
		fmt.Printf("Failed to Read: %v", err)
	}
	fmt.Println(val)
}
func NewFabricKVStoreInstance() (*BlockchainConnector.FabricConnector, error) {
	//var fabricConn *BlockchainConnector.FabricConnector

	channel1 := "kvchannel"

	chaincode1 := "kvstore"

	fmt.Println("Warning: Using default channel(kvchannel) and smartcontract(kvstore) ")

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"networks",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	defer gw.Close()
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %v", err)
		return nil, err
	}

	network1, err := gw.GetNetwork(channel1)
	if err != nil {
		fmt.Printf("Failed to get network: %v", err)
		return nil, err
	}
	kvcontract := network1.GetContract(chaincode1)
	result, err := kvcontract.SubmitTransaction("Set", "A13", "1300")
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}
	log.Println(string(result))

	return &BlockchainConnector.FabricConnector{KV: kvcontract}, nil
}

func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"..",
		"networks",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "User1@org1.example.com-cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put("appUser", identity)
}
