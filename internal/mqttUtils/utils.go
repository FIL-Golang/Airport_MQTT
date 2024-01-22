package mqttUtils

import (
	"Airport_MQTT/internal/config"
	"github.com/eclipse/paho.mqtt.golang"
	"log/slog"
)

func NewMqttClient() mqtt.Client {
	MqttConfig := config.GetMqttConfig()
	mqttClientOptions := mqtt.NewClientOptions()
	mqttClientOptions.AddBroker(MqttConfig.Broker.Url)
	mqttClientOptions.SetUsername(MqttConfig.Broker.Username)
	mqttClientOptions.SetPassword(MqttConfig.Broker.Password)
	mqttClientOptions.SetClientID(MqttConfig.Client.Id)
	mqttClient := mqtt.NewClient(mqttClientOptions)
	slog.Info("Connecting to MQTT broker")
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	slog.Info("Connected to MQTT broker")
	return mqttClient
}
