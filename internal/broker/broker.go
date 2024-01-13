package broker

import (
	"Airport_MQTT/internal/config/types"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

type Broker struct {
	client mqtt.Client
}

func NewBroker(cfg *types.ConfigFile) Broker {
	brokerURL := cfg.Mqtt.Url

	options := mqtt.NewClientOptions().AddBroker(brokerURL)
	client := mqtt.NewClient(options)
	if request := client.Connect(); request.Wait() && request.Error() != nil {
		fmt.Println("MQTT connexion error : ", request.Error())
		os.Exit(1)
	}
	return Broker{client: client}
}

func (broker Broker) SendMessage(topic, message string) {
	req := broker.client.Publish(topic, 1, false, message)
	req.Wait()
}
