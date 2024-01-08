package broker

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

type BrokerInterface interface {
}

type Broker struct {
	BrokerInterface
	BrokerAddress string
	BrokerPort    int
	QoS           byte
	ClientID      int
}

type BrokerConfig struct {
	client mqtt.Client
}

func NewBroker(brokerInterface BrokerInterface, brokerAddress string, brokerPort int, qos byte, clientID int) Broker {
	return Broker{
		BrokerAddress:   brokerAddress,
		BrokerPort:      brokerPort,
		QoS:             qos,
		ClientID:        clientID,
		BrokerInterface: brokerInterface,
	}
}

func ConfigureBroker(broker Broker) BrokerConfig {
	address := broker.BrokerAddress

	opts := mqtt.NewClientOptions().AddBroker(address)
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("MQTT connexion error : ", token.Error())
		os.Exit(1)
	}
	return BrokerConfig{client: client}
}

func (brokerConfig BrokerConfig) SendMessage(topic, message string) {
	req := brokerConfig.client.Publish(topic, 1, false, message)
	req.Wait()
}
