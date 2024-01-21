package main

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/database_recorder"
	"Airport_MQTT/internal/mqttUtils"
	"fmt"
)

const (
	SubscribeTopic = "/airports/+/sensors/+/+" // /airports/{airportIATA}/sensors/{sensorType}/{sensorId}
)

func init() {
	config.LoadConfig()
}

func main() {
	fmt.Println("Connecting to MQTT broker...")
	mqttClient := mqttUtils.NewMqttClient()
	fmt.Println("Connected to MQTT broker")

	fmt.Println("Connecting to MongoDB...")
	mqttHandler := database_recorder.NewDatabaseRecorderMqttHandler()
	fmt.Println("Connected to MongoDB")

	fmt.Println("Subscribing to topic...")
	mqttClient.Subscribe(SubscribeTopic, 0, mqttHandler.HandleValue)
	fmt.Println("Subscribed to topic : ", SubscribeTopic)

	select {}
}
