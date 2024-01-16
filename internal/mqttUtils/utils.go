package mqttUtils

import (
	"Airport_MQTT/internal/config"
	"github.com/eclipse/paho.mqtt.golang"
)

func NewMqttClient() mqtt.Client {
	MqttConfig := config.GetMqttConfig()
	mqttClientOptions := mqtt.NewClientOptions()
	mqttClientOptions.AddBroker(MqttConfig.Broker.Url)
	mqttClientOptions.SetUsername(MqttConfig.Broker.Username)
	mqttClientOptions.SetPassword(MqttConfig.Broker.Password)
	mqttClientOptions.SetClientID(MqttConfig.Client.Id)
	mqttClient := mqtt.NewClient(mqttClientOptions)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return mqttClient
}
