package main

import (
	"Airport_MQTT/internal/alert_manager"
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/mqttUtils"
	"log/slog"
)

const (
	SubscribeTopic = "/airports/+/sensors/+/+" // /airports/{airportIATA}/sensors/{sensorType}/{sensorId}
)

func init() {
	config.LoadConfig()
}

func main() {
	mqttClient := mqttUtils.NewMqttClient()

	mqttHandler := alert_manager.NewAlertManagerMqttHandler()

	slog.Info("Subscribing to topic: ", SubscribeTopic)
	mqttClient.Subscribe(SubscribeTopic, 0, mqttHandler.HandleValue)

	select {}
}
