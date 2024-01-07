package model

import "time"

const (
	Temperature = 0
	Pressure    = 1
	WindSpeed   = 2
)

type SensorData struct {
	SensorId  string  // format: <uuid>
	CodeIATA  string  // format: <3 letters>
	Nature    int     // 0: temperature, 1: pressure, 2: wind speed
	Value     float32 // value of the sensor
	Timestamp time.Time
}
