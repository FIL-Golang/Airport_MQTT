package main

import (
	"Airport_MQTT/internal/sensor"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"strconv"
)

type Config struct {
	SensorId    string `json:"sensorId"`
	IataCode    string `json:"iataCode"`
	Measurement string `json:"measurement"`
	Frequency   string `json:"frequency"`
}

func main() {
	configFilePath := flag.String("config", "config.json", "Path to the configuration file")
	flag.Parse()

	file, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		panic(err)
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	if _, err := uuid.Parse(config.SensorId); err != nil || config.SensorId == "" {
		config.SensorId = uuid.New().String()
		fmt.Println("SensorId was not an uuid new one is :", config.SensorId)
	}

	frequency, err := strconv.Atoi(config.Frequency)
	if err != nil {
		panic("Frequency must be an integer")
	}

	if config.IataCode == "" || config.Measurement == "" {
		fmt.Println("Missing required fields in config file")
		os.Exit(1)
	}

	mySensor := sensor.NewSensor(nil, config.SensorId, config.IataCode, config.Measurement, frequency)
	mySensor.StartSensor()
}
