package database_recorder

import (
	"Airport_MQTT/internal/persist"
	"Airport_MQTT/internal/utils"
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
	slog.Debug("Received message: " + string(msg.Payload()) + " on topic: " + msg.Topic())
	err, data := utils.Parse(msg)
	if err != nil {
		slog.Error("Error while parsing mqtt message: " + err.Error())
		return
	}

	slog.Info("Received data from sensor: " + data.SensorId + " from airport: " + data.AirportIATA + " of type: " + data.Nature.String())
	err = handler.repository.Store(data)
	if err != nil {
		slog.Error("Error while storing data: " + err.Error())
		return
	}
}
