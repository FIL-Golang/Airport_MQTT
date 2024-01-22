package database_recorder

import (
	"Airport_MQTT/internal/mqttUtils"
	"Airport_MQTT/internal/persist"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log/slog"
)

type DatabaseRecorderMqttHandler struct {
	repository persist.SensorDataRepository
}

func NewDatabaseRecorderMqttHandler() *DatabaseRecorderMqttHandler {
	repository := persist.NewSensorDataRepository()
	return &DatabaseRecorderMqttHandler{
		repository: repository,
	}
}

func (handler *DatabaseRecorderMqttHandler) HandleValue(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received value : %s on topic: %s\n", msg.Payload(), msg.Topic())
	err, data := mqttUtils.Parse(msg)
	if err != nil {
		slog.Error("Error parsing message: " + err.Error())
		return
	}
	err = handler.repository.Store(data)
	if err != nil {
		slog.Error("Error storing data: " + err.Error())
		return
	}
}
