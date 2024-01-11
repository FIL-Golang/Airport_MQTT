package sensor

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	err := godotenv.Load(filepath.Join(basepath, "../../.env"))
	if err != nil {
		fmt.Println("No .env file found")
	}
}

func TestFetchWeatherDataWithRealAPI(t *testing.T) {
	apiKey, exists := os.LookupEnv("WEATHER_API_KEY")
	if !exists || apiKey == "" {
		t.Fatalf("WEATHER_API_KEY is not set in .env file")
	}

	apiURL, exists := os.LookupEnv("WEATHER_URL")
	if !exists || apiURL == "" {
		t.Fatalf("WEATHER_URL is not set in .env file")
	}

	testSensor := Sensor{}
	city := "Paris"

	weatherData, err := testSensor.fetchWeatherData(city)
	if err != nil {
		t.Errorf("fetchWeatherData returned an error: %v", err)
	}
	if len(weatherData.Current) == 0 {
		t.Errorf("No data received from fetchWeatherData")
	}
}

func TestProcessWeatherData(t *testing.T) {
	testSensor := Sensor{}

	weatherData := WeatherResponse{
		Current: map[string]interface{}{
			"temperature": 23.5,
		},
	}

	measurement := testSensor.processWeatherData(weatherData, "temperature")
	if measurement.Value == nil || *measurement.Value != 23.5 {
		t.Errorf("Expected temperature 23.5, got %v", measurement.Value)
	}

	measurement = testSensor.processWeatherData(weatherData, "humidity")
	if measurement.Value != nil {
		t.Errorf("Expected nil value for unprovided measurement, got %v", measurement.Value)
	}
}

func TestGetDataApi(t *testing.T) {
	testSensor := Sensor{}

	city := "Paris"
	measurementType := "temperature"
	measurementData := testSensor.getDataApi(measurementType, city)

	if measurementData.Value == nil {
		t.Errorf("Expected non-nil value for measurement, got nil")
	}
	if measurementData.TypeMeasure != measurementType {
		t.Errorf("Expected measurement type %s, got %s", measurementType, measurementData.TypeMeasure)
	}
}
