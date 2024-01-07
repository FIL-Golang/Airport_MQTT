package sensor

import (
	"Airport_MQTT/api"
	"time"
)

type Sensor struct {
	SensorID    int
	ConfigID    int
	AirportCode string
	Measurement string
	Value       float64
	Timestamp   time.Time
	Frequency   int
}

func StartCaptureData(city string, measurement string, frequency int, dataChannel chan<- float64) {
	ticker := time.NewTicker(time.Second * time.Duration(frequency))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			measureValue, _, err := api.FetchWeatherData(city, measurement)
			if err != nil {
				continue
			}
			dataChannel <- float64(measureValue)
		}
	}
}
