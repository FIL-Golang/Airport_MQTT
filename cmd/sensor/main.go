package main

import (
	"Airport_MQTT/api"
	"Airport_MQTT/pkg/sensor"
	"encoding/json"
	"fmt"
	"github.com/krisukox/google-flights-api/iata"
	"io/ioutil"
	"time"
)

type Config struct {
	Sensors    []sensor.SensorData `json:"sensors"`
	MQTTConfig sensor.SensorConfig `json:"mqttConfig"`
}

func main() {
	configFile, err := ioutil.ReadFile("pkg/sensor/config.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier de configuration:", err)
		return
	}

	var config Config
	if err := json.Unmarshal(configFile, &config); err != nil {
		fmt.Println("Erreur lors du parsing du fichier de configuration:", err)
		return
	}

	for _, s := range config.Sensors {
		iataCode := s.AirportCode

		fmt.Println("test", iataCode)
		location := iata.IATATimeZone(iataCode)
		if location.City == "" {
			fmt.Println("Code IATA non supporté :", iataCode)
			continue
		}

		weatherData, err := api.FetchWeatherData(location.City)
		if err != nil {
			fmt.Println("Erreur lors de la récupération des données météo:", err)
			continue
		}

		fmt.Printf("Météo pour %s: Température %d°C, Vitesse du vent %d km/h, Humidité %d%%\n",
			location.City, weatherData.Current.Temperature, weatherData.Current.WindSpeed, weatherData.Current.Humidity)

		s.Value = float64(weatherData.Current.Temperature)
		s.Timestamp = time.Now()

		s.Display()
	}

	config.MQTTConfig.Display()
}
