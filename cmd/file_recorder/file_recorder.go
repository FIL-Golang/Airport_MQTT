package main

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/file_recorder"
	"Airport_MQTT/internal/utils"
	"fmt"
)

const (
	SubscribeTopic = "/airports/+/sensors/+/+" // /airports/{airportIATA}/sensors/{sensorType}/{sensorId}
)

func init() {
	config.LoadConfigFromArgs()
}

func main() {
	fmt.Println("Connecting to MQTT broker...")
	mqttClient := utils.NewMqttClient()
	fmt.Println("Connected to MQTT broker")

	fmt.Println("Connecting to MongoDB...")
	mqttHandler := file_recorder.NewFileRecoderMqttHandler()
	fmt.Println("Connected to MongoDB")

	fmt.Println("Subscribing to topic...")
	mqttClient.Subscribe(SubscribeTopic, 0, mqttHandler.HandleValue)
	fmt.Println("Subscribed to topic : ", SubscribeTopic)

	select {}

}
