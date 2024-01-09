package mqttUtils

import (
	"Airport_MQTT/internal/persist"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type FileRecoderMqttHandler struct {
	recorder persist.SensorFileRecorder
	parser   Parser
}

func NewFileRecoderMqttHandler() *FileRecoderMqttHandler {
	recorder := persist.NewSensorFileRecorder()
	return &FileRecoderMqttHandler{
		recorder: recorder,
		parser:   NewParser(),
	}
}

func (this *FileRecoderMqttHandler) HandleValue(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received value : %s on topic: %s\n", msg.Payload(), msg.Topic())
	data := this.parser.Parse(msg)
	data.Timestamp = time.Now()
	_, err := this.recorder.Store(data)
	if err != nil {
		println(err.Error())
	}
}
