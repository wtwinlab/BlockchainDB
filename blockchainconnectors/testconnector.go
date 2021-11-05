package connectors

type Testconnector struct {
}

func (t *Testconnector) Read(key string) (string, error) {
	return "", nil
}

func (t *Testconnector) Write(key, value string) (string, error) {
	return "", nil
}
