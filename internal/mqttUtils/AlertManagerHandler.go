package mqttUtils

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/config/types"
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
	data := this.parser.Parse(msg)

	conf := config.LoadConfig(&types.AlertConfigFile{}, "alerts.yaml").(*types.AlertConfigFile)

	for _, alert := range conf.Alerts {
		if alert.Airport == data.CodeIATA {
			if data.Nature == model.Temperature {
				if data.Value > float32(alert.Temperature) {
					this.publishAlert(client, data, data.Nature, "Alerte température")
					fmt.Println("Alerte température")
				}
			} else if data.Nature == model.Pressure {
				if data.Value > float32(alert.Pressure) {
					this.publishAlert(client, data, data.Nature, "Alerte pression")
					fmt.Println("Alerte pression")
				}
			} else if data.Nature == model.WindSpeed {
				if data.Value > float32(alert.Wind) {
					this.publishAlert(client, data, data.Nature, "Alerte vitesse du vent")
					fmt.Println("Alerte vitesse du vent")
				}
			}
		}
	}
}

func (h *AlertManagerMqttHandler) publishAlert(client mqtt.Client, data model.SensorData, alertType int, message string) {
	client.Publish("/airports/"+data.CodeIATA+"/alerts/"+model.SensorNatureFromInt(alertType)+"/"+data.SensorId, 0, false, message)
}
