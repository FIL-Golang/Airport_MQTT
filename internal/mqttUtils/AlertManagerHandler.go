package mqttUtils

import (
	"Airport_MQTT/internal/model"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Alert struct {
	Name string
	Max  float32
}

type AlertManagerMqttHandler struct {
	parser Parser
	alerts map[string][]Alert
}

func NewAlertManagerMqttHandler() *AlertManagerMqttHandler {
	return &AlertManagerMqttHandler{
		parser: NewParser(),
		alerts: make(map[string][]Alert),
	}
}

func (this *AlertManagerMqttHandler) HandleValue(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received value : %s on topic: %s\n", msg.Payload(), msg.Topic())
	data := this.parser.Parse(msg)

	if data.Nature == model.Temperature {
		if data.Value > 10 {
			fmt.Println("Alerte température")
		}
	} else if data.Nature == model.Pressure {
		if data.Value > 1013 {
			fmt.Println("Alerte pression")
		}
	} else if data.Nature == model.WindSpeed {
		if data.Value > 100 {
			fmt.Println("Alerte vitesse du vent")
		}
	}
}
