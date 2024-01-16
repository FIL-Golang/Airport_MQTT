package model

import (
	"fmt"
	"time"
)

type Nature int

const (
	Undefined Nature = iota
	Temperature
	WindSpeed
	Pressure
)

// SensorData is the data sent by the sensors to the broker
type SensorData struct {
	SensorId    string  // format: <uuid>
	AirportIATA string  // format: <3 letters>
	Nature      Nature  // 0: temperature, 1: pressure, 2: wind speed
	Value       float32 // value of the sensor
	Timestamp   time.Time
}

func (nature Nature) String() string {
	return [...]string{"undefined", "temperature", "wind_speed", "pressure"}[nature]
}

// MarshalJSON marshals the enum as a quoted json string
// Permit to send the string instead of the int when using json.Marshal
func (nature Nature) MarshalJSON() ([]byte, error) {
	return []byte(`"` + nature.String() + `"`), nil
}

func SensorNatureFromString(nature string) Nature {
	fmt.Println("PARSING NATURE: ", nature)
	switch nature {
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
	Type        Nature    `bson:"sensorType" json:"sensorType"`
	Readings    []Reading `bson:"readings" json:"readings,omitempty"`
}

type Reading struct {
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Value     float32   `bson:"value" json:"value"`
}

type Average struct {
	Avg        float64 `bson:"avg"`
	SensorType Nature  `bson:"sensorType"`
}
