package connectors

type BlockchainConnector interface {
	Get(string) (string, error)
	Put(key, value string) error
}
