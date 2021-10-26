package connectors

type BlockchainConnector interface {
	Get() (string, error)
	Put(key, value string) error
}
