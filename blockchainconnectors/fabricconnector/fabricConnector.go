package connector

import (
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type FabricConnector struct {
	KV *gateway.Contract
}

func (fabconn *FabricConnector) Read(key string) (string, error) {
	//log.Println("--> Submit Transaction: InitLedger, function get value of key on the ledger")
	result, err := fabconn.KV.EvaluateTransaction("Get", key)
	if err != nil {
		log.Printf("Failed to Submit transaction: %v", err)
		return "", err
	}
	return string(result), nil

}
func (fabconn *FabricConnector) Write(key string, value string) error {
	//log.Println("--> Submit Transaction: InitLedger, function set value of key on the ledger")
	_, err := fabconn.KV.SubmitTransaction("Set", key, value)
	if err != nil {
		log.Printf("Failed to Submit transaction: %v", err)
	}
	return err
}
