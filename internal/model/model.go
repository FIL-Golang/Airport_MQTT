package model

import "time"

type Nature int

const (
	Temperature Nature = iota
	WindSpeed
	Pressure
)

// SensorData is the data sent by the sensors to the broker
type SensorData struct {
	SensorId    string  // format: <uuid>
	AirportIATA string  // format: <3 letters>
	Nature      Nature  // 0: temperature, 1: pressure, 2: wind speed
	Value       float64 // value of the sensor
	Timestamp   string
}

// Sensor is the data that we return from the API
type Sensor struct {
	SensorId    string
	AirportIATA string
	Type        Nature
	Readings    []Reading
}

type Reading struct {
	Timestamp time.Time
	Value     float32
}

func (nature Nature) String() string {
	return [...]string{"temperature", "wind_speed", "pressure"}[nature]
}

func SensorNatureFromString(nature string) Nature {
	switch nature {
	case "temperature":
		return Temperature
	case "pressure":
		return Pressure
	case "wind_speed":
		return WindSpeed
	default:
		return -1
	}
}
