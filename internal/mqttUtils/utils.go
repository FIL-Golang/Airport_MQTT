package mqttUtils

import (
	"Airport_MQTT/internal/config"
	"github.com/eclipse/paho.mqtt.golang"
)

func NewMqttClient() mqtt.Client {
	conf := config.LoadConfig()
	mqttClientOptions := mqtt.NewClientOptions()
	mqttClientOptions.AddBroker(conf.Mqtt.Url)
	mqttClientOptions.SetUsername(conf.Mqtt.Username)
	mqttClientOptions.SetPassword(conf.Mqtt.Password)
	mqttClientOptions.SetClientID(conf.Name)
	return mqtt.NewClient(mqttClientOptions)
}
