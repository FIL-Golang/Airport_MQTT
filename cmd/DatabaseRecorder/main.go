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

	fmt.Println("Connecting to MongoDB...")
	mqttHandler := mqttUtils.NewDatabaseRecorderMqttHandler()
	fmt.Println("Connected to MongoDB")

	fmt.Println("Subscribing to topic...")
	mqttClient.Subscribe("/airports/+/+/+", 0, mqttHandler.HandleValue)
	fmt.Println("Subscribed to topic airports/+/+/+")

	select {}

}
