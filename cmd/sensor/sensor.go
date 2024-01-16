package main

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/sensor"
	"fmt"
	"os"
)

func init() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: database_recorder <config_file>")
		os.Exit(1)
	}
	config.LoadConfig(args[1])
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
