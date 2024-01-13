package main

import (
	"Airport_MQTT/internal/sensor"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	SensorId    string `yaml:"sensorId"`
	IataCode    string `yaml:"iataCode"`
	Measurement string `yaml:"measurement"`
	Frequency   int    `yaml:"frequency"`
}

func main() {
	config := loadConfig()
	validateConfig(&config)
	mySensor := createSensor(config)
	mySensor.StartSensor()
}

func loadConfig() Config {
	configFilePath := flag.String("config", "config.yaml", "Path to the configuration file")
	flag.Parse()

	file, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func validateConfig(config *Config) error {
	if _, err := uuid.Parse(config.SensorId); err != nil || config.SensorId == "" {
		config.SensorId = uuid.New().String()
		fmt.Println("SensorId was not an uuid new one is :", config.SensorId)
	}

	if config.IataCode == "" || config.Measurement == "" {
		fmt.Println("Missing required fields in config file")
		os.Exit(1)
	}
	return nil
}

func createSensor(config Config) *sensor.Sensor {
	s := sensor.NewSensor(nil, config.SensorId, config.IataCode, config.Measurement, config.Frequency)
	return &s
}
