package connectors

import (
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type FabricConnector struct {
	KV *gateway.Contract
}

func (fabconn *FabricConnector) Read(key string) (string, error) {
	log.Println("--> Submit Transaction: InitLedger, function creates the initial set of accounts on the ledger")
	result, err := fabconn.KV.EvaluateTransaction("Get", key)
	if err != nil {
		log.Printf("Failed to Submit transaction: %v", err)
		return "", err
	}
	return string(result), nil

}
func (fabconn *FabricConnector) Write(key string, value string) (string, error) {
	result, err := fabconn.KV.SubmitTransaction("Set", key, value)
	if err != nil {
		log.Printf("Failed to Submit transaction: %v", err)
		return "", err
	}
	return string(result), nil
}
