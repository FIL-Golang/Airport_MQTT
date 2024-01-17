package main

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/database_recorder"
	"Airport_MQTT/internal/utils"
	"log/slog"
)

const (
	SubscribeTopic = "/airports/+/sensors/+/+" // /airports/{airportIATA}/sensors/{sensorType}/{sensorId}
)

func init() {
	config.LoadConfigFromArgs()
}

func main() {
	slog.Info("Starting " + config.GetExeName())
	mqttClient := utils.NewMqttClient()
	mqttHandler := database_recorder.NewDatabaseRecorderMqttHandler()
	slog.Info("Subscribing to " + SubscribeTopic)
	mqttClient.Subscribe(SubscribeTopic, 0, mqttHandler.HandleValue)

	select {}
}
