package alert_manager

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/model"
	"Airport_MQTT/internal/mqttUtils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log/slog"
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
		slog.Error("Error parsing message: ", err)
		return
	}

	alerts := config.GetAlerts()

	for _, alert := range alerts {
		if alert.Airport == data.AirportIATA {
			if data.Type == model.Temperature {
				if data.Value > float32(alert.Temperature) {
					this.publishAlert(client, data, "Alerte température")
					slog.Info("Alerte température")
				}
			} else if data.Type == model.Pressure {
				if data.Value > float32(alert.Pressure) {
					this.publishAlert(client, data, "Alerte pression")
					slog.Info("Alerte pression")
				}
			} else if data.Type == model.WindSpeed {
				if data.Value > float32(alert.Wind) {
					this.publishAlert(client, data, "Alerte vitesse du vent")
					slog.Info("Alerte vitesse du vent")
				}
			}
		}
	}
}

func (h *AlertManagerMqttHandler) publishAlert(client mqtt.Client, data model.SensorData, message string) {
	client.Publish(mqttUtils.GetAlertsTopic(data), byte(config.GetMqttConfig().Client.QOS), false, message)
}
