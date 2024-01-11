package main

import (
	"Airport_MQTT/internal/sensor"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
)

func main() {
	sensorIdFlag := flag.String("sensorId", "", "ID of the sensor")
	iataCodeFlag := flag.String("iataCode", "", "IATA code")
	measurementFlag := flag.String("measurement", "", "Type of measurement")
	frequencyStrFlag := flag.String("frequency", "", "Frequency of measurement")

	flag.Parse()

	if _, err := uuid.Parse(*sensorIdFlag); err != nil || *sensorIdFlag == "" {
		*sensorIdFlag = uuid.New().String()
		fmt.Println("SensorId was not an uuid new one is :", *sensorIdFlag)
	}

	frequency, err := strconv.Atoi(*frequencyStrFlag)
	if err != nil {
		panic("Frequency must be an integer")
	}

	if *iataCodeFlag == "" || *measurementFlag == "" {
		fmt.Println("Missing required flags")
		flag.Usage()
		os.Exit(1)
	}

	mySensor := sensor.NewSensor(nil, *sensorIdFlag, *iataCodeFlag, *measurementFlag, frequency)
	mySensor.StartSensor()
}
