package alert_manager

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/model"
	"Airport_MQTT/internal/mqttUtils"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type AlertManagerMqttHandler struct {
	alerts map[string][]config.Alert
}

func NewAlertManagerMqttHandler() *AlertManagerMqttHandler {
	return &AlertManagerMqttHandler{
		alerts: make(map[string][]config.Alert),
	}
}

func (this *AlertManagerMqttHandler) HandleValue(client mqtt.Client, msg mqtt.Message) {
	err, data := mqttUtils.Parse(msg)
	if err != nil {
		println(err.Error())
		return
	}

	alerts := config.GetAlerts()

	for _, alert := range alerts {
		if alert.Airport == data.AirportIATA {
			if data.Nature == model.Temperature {
				if data.Value > float32(alert.Temperature) {
					this.publishAlert(client, data, "Alerte température")
					fmt.Println("Alerte température")
				}
			} else if data.Nature == model.Pressure {
				if data.Value > float32(alert.Pressure) {
					this.publishAlert(client, data, "Alerte pression")
					fmt.Println("Alerte pression")
				}
			} else if data.Nature == model.WindSpeed {
				if data.Value > float32(alert.Wind) {
					this.publishAlert(client, data, "Alerte vitesse du vent")
					fmt.Println("Alerte vitesse du vent")
				}
			}
		}
	}
}

func (h *AlertManagerMqttHandler) publishAlert(client mqtt.Client, data model.SensorData, message string) {
	client.Publish(mqttUtils.GetAlertsTopic(data), 0, false, message)
}
