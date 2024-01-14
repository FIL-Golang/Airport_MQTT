package sensor

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WeatherResponse struct {
	Location map[string]interface{} `json:"location"`
	Current  map[string]interface{} `json:"current"`
}

func (s *Sensor) fetchWeatherData(city string) (WeatherResponse, error) {
	apiConfig := config.GetApiConfig()
	url := fmt.Sprintf(apiConfig.Url, apiConfig.SecretKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var weatherData WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		return WeatherResponse{}, err
	}

	return weatherData, nil
}

func (s *Sensor) processWeatherData(weatherData WeatherResponse, measurement string) model.SensorData {
	measureValue, ok := weatherData.Current[measurement].(float64)
	if !ok {
		return model.SensorData{}
	}
	return model.SensorData{
		SensorId:    s.DeviceId,
		AirportIATA: s.AirportIATA,
		Nature:      model.SensorNatureFromString(s.Type),
		Value:       measureValue,
		Timestamp:   time.Now().Format("2006-01-02-15-04-05"),
	}
}

func (s *Sensor) getDataApi(measurement string, city string) model.SensorData {
	weatherData, err := s.fetchWeatherData(city)
	if err != nil {
		return model.SensorData{}
	}
	return s.processWeatherData(weatherData, measurement)
}
