package model

import (
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

func (typ Type) String() string {
	return [...]string{"undefined", "temperature", "wind_speed", "pressure"}[typ]
}

// MarshalJSON marshals the enum as a quoted json string
// Permit to send the string instead of the int when using json.Marshal
func (typ Type) MarshalJSON() ([]byte, error) {
	return []byte(`"` + typ.String() + `"`), nil
}

func SensorTypeFromString(typ string) Type {
	switch typ {
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
