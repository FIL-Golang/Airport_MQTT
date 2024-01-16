package file_recorder

import (
	"Airport_MQTT/internal/mqttUtils"
	"Airport_MQTT/internal/persist"
	"fmt"

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
	fmt.Printf("Received value : %s on topic: %s\n", msg.Payload(), msg.Topic())
	err, data := mqttUtils.Parse(msg)
	err = handler.recorder.Store(data)
	if err != nil {
		println(err.Error())
	}
}