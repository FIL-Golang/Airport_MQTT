package sensor

import "fmt"

type Config struct {
	BrokerAddress string
	BrokerPort    int
	QoS           byte
	ClientID      int
}

func NewConfig(brokerAddress string, brokerPort int, qos byte, clientID int) *Config {
	return &Config{
		BrokerAddress: brokerAddress,
		BrokerPort:    brokerPort,
		QoS:           qos,
		ClientID:      clientID,
	}
}

func (sc *Config) Display() {
	fmt.Printf("Broker Address: %s\n", sc.BrokerAddress)
	fmt.Printf("Broker Port: %d\n", sc.BrokerPort)
	fmt.Printf("QoS: %d\n", sc.QoS)
	fmt.Printf("Client ID: %s\n", sc.ClientID)
}
