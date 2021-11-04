package connectors

type BlockchainConnector interface {
	Read(string) (string, error)
	Write(string, string) error
}
