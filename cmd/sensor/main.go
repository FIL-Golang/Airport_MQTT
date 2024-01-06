package main

import (
	"Airport_MQTT/pkg/sensor"
	"time"
)

func main() {
	sensorData := sensor.NewSensorData(123, "CDG", "Temperature", 26.5, time.Now())
	sensorConfig := sensor.NewSensorConfig("mqtt.com", 1883, 0, 123)

	sensorData.Display()
	sensorConfig.Display()
}
