package mqttUtils

import (
	"Airport_MQTT/internal/persist"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type DatabaseRecorderMqttHandler struct {
	repository persist.SensorDataRepository
	parser     Parser
}

func NewDatabaseRecorderMqttHandler() *DatabaseRecorderMqttHandler {
	repository := persist.NewSensorDataRepository()
	return &DatabaseRecorderMqttHandler{
		repository: repository,
		parser:     NewParser(),
	}
}

func (this *DatabaseRecorderMqttHandler) HandleValue(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received value : %s on topic: %s\n", msg.Payload(), msg.Topic())
	data := this.parser.Parse(msg)
	data.Timestamp = time.Now()
	_, err := this.repository.Store(data)
	if err != nil {
		println(err.Error())
	}
}
