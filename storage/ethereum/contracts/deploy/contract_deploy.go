package main

import (
	"fmt"
	"log"

	ClientSDK "github.com/sbip-sg/BlockchainDB/storage/ethereum/clientSDK"
)

func main() {

	ethnode := "http://localhost:8000"
	hexkey := "2243d6f09ad779cc50ccecc3a6779b62ba6c0bb1390e6b0dfcd8fca5eea98193"

	address, tx, instance, err := ClientSDK.DeployEthereumKVStoreContract(ethnode, hexkey)
	if err != nil {
		log.Fatal("DeployEthereumKVStoreContract", err)
	}

	fmt.Println(address)
	fmt.Println(tx)
	_ = instance

}
