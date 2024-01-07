package sensor

type Config struct {
	BrokerAddress string
	BrokerPort    int
	QoS           byte
	ClientID      int
}
