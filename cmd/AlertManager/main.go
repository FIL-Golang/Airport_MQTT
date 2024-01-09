package main

import (
	"Airport_MQTT/internal/mqttUtils"
	"fmt"
)

func main() {
	fmt.Println("Connecting to MQTT broker...")
	mqttClient := mqttUtils.NewMqttClient()
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Connected to MQTT broker")

	mqttHandler := mqttUtils.NewAlertManagerMqttHandler()

	fmt.Println("Subscribing to sensors topics...")
	mqttClient.Subscribe("/airports/+/sensors/+/+", 0, mqttHandler.HandleValue)
	fmt.Println("Subscribed to sensors topics airports/+/sensors/+/+")

	select {}
}
