package main

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/sensor"
)

func init() {
	config.LoadConfigFromArgs()
}

func main() {
	mySensor := createSensor()
	mySensor.StartSensor()
}

func createSensor() *sensor.Sensor {
	s := sensor.NewSensor(nil,
		config.GetSensorConfig().DeviceId,
		config.GetSensorConfig().AirportIATA,
		config.GetSensorConfig().SensorType,
		config.GetSensorConfig().Frequency)
	return &s
}
