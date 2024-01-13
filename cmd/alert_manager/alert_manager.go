package main

import (
	"Airport_MQTT/internal/alert_manager"
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/mqttUtils"
	"fmt"
	"os"
)

const (
	SubscribeTopic = "/airports/+/sensors/+/+" // /airports/{airportIATA}/sensors/{sensorType}/{sensorId}
)

func init() { // TODO : refactor to reuse in each executable (or in config.LoadConfig)
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: alert_manager <config_file>")
		os.Exit(1)
	}
	config.LoadConfig(args[1])
}

func main() {
	fmt.Println("Connecting to MQTT broker...")
	mqttClient := mqttUtils.NewMqttClient()
	fmt.Println("Connected to MQTT broker")

	mqttHandler := alert_manager.NewAlertManagerMqttHandler()

	fmt.Println("Subscribing to sensors topics...")
	mqttClient.Subscribe(SubscribeTopic, 0, mqttHandler.HandleValue)
	fmt.Println("Subscribed to sensors topics : ", SubscribeTopic)

	select {}
}
