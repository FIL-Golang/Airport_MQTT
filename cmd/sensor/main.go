package main

import (
	"Airport_MQTT/api"
	"Airport_MQTT/pkg/sensor"
	"fmt"
	"github.com/krisukox/google-flights-api/iata"
	"time"
)

func main() {
	iataCode := "CDG"

	location := iata.IATATimeZone(iataCode)
	if location.City == "" {
		fmt.Println("Code IATA non supporté :", iataCode)
		return
	}

	fmt.Printf("Aéroport : %s\n", location.City)

	weatherData, err := api.FetchWeatherData(location.City)
	if err != nil {
		fmt.Println("Erreur lors de la récupération des données météo:", err)
		return
	}

	fmt.Printf("Météo pour %s: Température %d°C, Vitesse du vent %d km/h, Humidité %d%%\n",
		location.City, weatherData.Current.Temperature, weatherData.Current.WindSpeed, weatherData.Current.Humidity)

	sensorData := sensor.NewSensorData(123, iataCode, "Temperature", float64(weatherData.Current.Temperature), time.Now())
	sensorConfig := sensor.NewSensorConfig("mqtt.com", 1883, 0, 123)

	sensorData.Display()
	sensorConfig.Display()
}
