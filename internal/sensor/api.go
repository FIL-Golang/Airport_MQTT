package sensor

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type WeatherResponse struct {
	Location map[string]interface{} `json:"location"`
	Current  map[string]interface{} `json:"current"`
}

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	err := godotenv.Load(filepath.Join(basepath, "../../.env"))
	if err != nil {
		fmt.Println("No .env file found")
	}
}

func (s *Sensor) fetchWeatherData(city string) (WeatherResponse, error) {
	apiKey, _ := os.LookupEnv("WEATHER_API_KEY")
	apiUrl, _ := os.LookupEnv("WEATHER_URL")
	url := fmt.Sprintf(apiUrl, apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var weatherData WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		return WeatherResponse{}, err
	}

	return weatherData, nil
}

func (s *Sensor) processWeatherData(weatherData WeatherResponse, measurement string) MeasurementData {
	measureValue, ok := weatherData.Current[measurement]
	if !ok {
		return MeasurementData{TypeMeasure: measurement, Value: nil, Timestamp: time.Now().Format(time.RFC3339)}
	}

	value, isFloat := measureValue.(float64)
	if !isFloat {
		return MeasurementData{TypeMeasure: measurement, Value: nil, Timestamp: time.Now().Format(time.RFC3339)}
	}

	return MeasurementData{TypeMeasure: measurement, Value: &value, Timestamp: time.Now().Format(time.RFC3339)}
}

func (s *Sensor) getDataApi(measurement string, city string) MeasurementData {
	weatherData, err := s.fetchWeatherData(city)
	if err != nil {
		return MeasurementData{TypeMeasure: measurement, Value: nil, Timestamp: time.Now().Format(time.RFC3339)}
	}

	return s.processWeatherData(weatherData, measurement)
}
