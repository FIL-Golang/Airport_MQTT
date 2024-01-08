package main

import (
	"Airport_MQTT/internal/sensor"
)

func main() {
	/*if len(os.Args) < 5 {
		panic("Usage: main <sensorId> <iataCode> <measurement> <frequency>")
	}

	sensorId := os.Args[1]
	iataCode := os.Args[2]
	measurement := os.Args[3]
	frequencyStr := os.Args[4]

	frequency, err := strconv.Atoi(frequencyStr)
	if err != nil {
		panic("Frequency must be an integer")
	}

	mySensor := sensor.NewSensor(nil, sensorId, iataCode, measurement, frequency)*/
	mySensor := sensor.NewSensor(nil, "1", "CDG", "wind_speed", 2)

	mySensor.StartSensor()
}
