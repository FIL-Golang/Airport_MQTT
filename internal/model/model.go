package model

import (
	"fmt"
	"time"
)

type Type int

const (
	Undefined Type = iota
	Temperature
	WindSpeed
	Pressure
)

// SensorData is the data sent by the sensors to the broker
type SensorData struct {
	SensorId    string  // format: <uuid>
	AirportIATA string  // format: <3 letters>
	Type        Type    // 0: temperature, 1: pressure, 2: wind speed
	Value       float32 // value of the sensor
	Timestamp   time.Time
}

func (_type Type) String() string {
	return [...]string{"undefined", "temperature", "wind_speed", "pressure"}[_type]
}

// MarshalJSON marshals the enum as a quoted json string
// Permit to send the string instead of the int when using json.Marshal
func (_type Type) MarshalJSON() ([]byte, error) {
	return []byte(`"` + _type.String() + `"`), nil
}

func SensorTypeFromString(_type string) Type {
	fmt.Println("PARSING TYPE: ", _type)
	switch _type {
	case "temperature":
		return Temperature
	case "pressure":
		return Pressure
	case "wind_speed":
		return WindSpeed
	default: // undefined
		return Undefined
	}
}

//Types for fetching data from the database

type Sensor struct {
	SensorId    string    `bson:"sensorId" json:"sensorId"`
	AirportIATA string    `bson:"airportIATA" json:"airportIATA"`
	Type        Type      `bson:"sensorType" json:"sensorType"`
	Readings    []Reading `bson:"readings" json:"readings,omitempty"`
}

type Reading struct {
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Value     float32   `bson:"value" json:"value"`
}

type Average struct {
	Avg        float64 `bson:"avg"`
	SensorType Type    `bson:"sensorType"`
}
