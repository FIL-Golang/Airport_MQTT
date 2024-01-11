package main

import (
	"Airport_MQTT/internal/sensor"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
)

type Args struct {
	SensorId    string
	IataCode    string
	Measurement string
	Frequency   string
}

func parseArgs(args Args) (string, string, string, int, error) {
	var sensorId string
	if _, err := uuid.Parse(args.SensorId); err != nil || args.SensorId == "" {
		sensorId = uuid.New().String()
		fmt.Println("SensorId was not an uuid new one is:", sensorId)
	} else {
		sensorId = args.SensorId
	}

	frequency, err := strconv.Atoi(args.Frequency)
	if err != nil {
		return "", "", "", 0, fmt.Errorf("frequency must be an integer")
	}

	if args.IataCode == "" || args.Measurement == "" {
		return "", "", "", 0, fmt.Errorf("missing required flags")
	}

	return sensorId, args.IataCode, args.Measurement, frequency, nil
}

func main() {
	sensorIdFlag := flag.String("sensorId", "", "ID of the sensor")
	iataCodeFlag := flag.String("iataCode", "", "IATA code")
	measurementFlag := flag.String("measurement", "", "Type of measurement")
	frequencyStrFlag := flag.String("frequency", "", "Frequency of measurement")

	flag.Parse()

	args := Args{
		SensorId:    *sensorIdFlag,
		IataCode:    *iataCodeFlag,
		Measurement: *measurementFlag,
		Frequency:   *frequencyStrFlag,
	}

	sensorId, iataCode, measurement, frequency, err := parseArgs(args)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(1)
	}

	mySensor := sensor.NewSensor(nil, sensorId, iataCode, measurement, frequency)
	mySensor.StartSensor()
}
