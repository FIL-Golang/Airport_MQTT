package file_recorder

import (
	"Airport_MQTT/internal/mqttUtils"
	"Airport_MQTT/internal/persist"
	"log/slog"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type FileRecoderMqttHandler struct {
	recorder persist.SensorFileRecorder
}

func NewFileRecoderMqttHandler() *FileRecoderMqttHandler {
	recorder := persist.NewSensorFileRecorder()
	return &FileRecoderMqttHandler{
		recorder: recorder,
	}
}

func (handler *FileRecoderMqttHandler) HandleValue(client mqtt.Client, msg mqtt.Message) {
	slog.Info("Received value : " + string(msg.Payload()) + " on topic: " + msg.Topic())
	err, data := mqttUtils.Parse(msg)
	if err != nil {
		slog.Error("Error while parsing data: " + err.Error())
		return
	}
	err = handler.recorder.Store(data)
	if err != nil {
		slog.Error("Error while storing data: " + err.Error())
	}
}
