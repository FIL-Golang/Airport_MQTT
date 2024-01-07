package api

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type WeatherResponse struct {
	Location map[string]interface{} `json:"location"`
	Current  map[string]interface{} `json:"current"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func FetchWeatherData(city string, measurement string) (int, string, error) {
	apiKey, exists := os.LookupEnv("WEATHER_API_KEY")
	if !exists {
		return 0, "", fmt.Errorf("API key not set in environment")
	}

	apiUrl, exists := os.LookupEnv("WEATHER_URL")
	if !exists {
		return 0, "", fmt.Errorf("API URL not set in environment")
	}

	url := fmt.Sprintf(apiUrl, apiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	var weatherData WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return 0, "", err
	}

	measureValue, ok := weatherData.Current[measurement]
	if !ok {
		return 0, "", fmt.Errorf("measurement not found")
	}

	measureValueInt, ok := measureValue.(float64)
	if !ok {
		return 0, "", fmt.Errorf("measurement is not a float64")
	}

	localTime, ok := weatherData.Location["localtime"]
	if !ok {
		return 0, "", fmt.Errorf("localtime not found")
	}

	localTimeString, ok := localTime.(string)
	if !ok {
		return 0, "", fmt.Errorf("localtime is not a string")
	}

	return int(measureValueInt), localTimeString, nil
}
