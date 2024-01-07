package main

import (
	"Airport_MQTT/internal/mqttUtils"
	"Airport_MQTT/internal/persist"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

func main() {
	fmt.Println("Connecting to MQTT broker...")
	mqttClient := mqttUtils.NewMqttClient()
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Connected to MQTT broker")

	fmt.Println("Connecting to MongoDB...")
	mqttHandler := newSensorDataMqttHandler()
	fmt.Println("Connected to MongoDB")

	fmt.Println("Subscribing to topic...")
	mqttClient.Subscribe("/airports/+/+/+", 0, mqttHandler.handleValue)
	fmt.Println("Subscribed to topic airports/+/+/+")

	select {}

}

type SensorDataMqttHandler struct {
	repository persist.SensorDataRepository
	parser     mqttUtils.Parser
}

func newSensorDataMqttHandler() *SensorDataMqttHandler {
	repository := persist.NewSensorDataRepository()
	return &SensorDataMqttHandler{
		repository: repository,
		parser:     mqttUtils.NewParser(),
	}
}

func (this *SensorDataMqttHandler) handleValue(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received value : %s on topic: %s\n", msg.Payload(), msg.Topic())
	data := this.parser.Parse(msg)
	data.Timestamp = time.Now()
	_, err := this.repository.Store(data)
	if err != nil {
		println(err.Error())
	}
}
