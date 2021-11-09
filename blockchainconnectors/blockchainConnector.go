package connectors

type BlockchainConnector interface {
	Read(string) (string, error)
	Write(string, string) (string, error)
	Verify(string, string) (bool, error)
}
