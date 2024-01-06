package sensor

import (
	"fmt"
	"time"
)

type SensorData struct {
	SensorID    int
	AirportCode string
	Measurement string
	Value       float64
	Timestamp   time.Time
}

func NewSensorData(sensorID int, airportCode string, measurement string, value float64, timestamp time.Time) *SensorData {
	return &SensorData{
		SensorID:    sensorID,
		AirportCode: airportCode,
		Measurement: measurement,
		Value:       value,
		Timestamp:   timestamp,
	}
}

func (sd *SensorData) Display() {
	fmt.Printf("Sensor ID: %d\n", sd.SensorID)
	fmt.Printf("Airport Code: %s\n", sd.AirportCode)
	fmt.Printf("Measurement: %s\n", sd.Measurement)
	fmt.Printf("Value: %.2f\n", sd.Value)
	fmt.Printf("Timestamp: %s\n", sd.Timestamp.Format("2006-01-02 15:04:05"))
}
