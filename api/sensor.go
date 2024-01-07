package api

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Location struct {
		LocalTime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		Temperature int `json:"temperature"`
		WindSpeed   int `json:"wind_speed"`
		Humidity    int `json:"humidity"`
	} `json:"current"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func FetchWeatherData(city string) (*WeatherResponse, error) {
	apiKey, exists := os.LookupEnv("WEATHER_API_KEY")
	apiUrl, exists := os.LookupEnv("WEATHER_URL")
	if !exists {
		return nil, fmt.Errorf("API key not set in environment")
	}

	url := fmt.Sprintf(apiUrl, apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherResponse
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return nil, err
	}

	return &weatherData, nil
}
