package api

import (
	"Airport_MQTT/internal/model"
	"Airport_MQTT/internal/persist"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

//Example of sending data

func SendDataExample(w http.ResponseWriter, r *http.Request) {
	sensorData := model.SensorData{
		SensorId:    "1",
		AirportIATA: "DBX",
		Nature:      model.Temperature, Value: rand.Float32(),
		Timestamp: time.Now().AddDate(0, 0, -1)}

	err := persist.NewSensorDataRepository().Store(sensorData)
	if err != nil {
		fmt.Println(w, "Error sending data: %v", err)
		return
	}
}

//Example of sending data

func GetSensor(w http.ResponseWriter, r *http.Request) {
	sensorID := r.URL.Query().Get("sensorID")
	airportIATA := r.URL.Query().Get("airportIATA")
	_type := r.URL.Query().Get("type")

	filter := persist.Filter{
		From: time.Unix(0, 0),
		To:   time.Now(),
	}

	if sensorID != "" {
		filter.SensorId = sensorID
	}
	if airportIATA != "" {
		filter.AirportIATA = airportIATA
	}
	if _type != "" {
		filter.Type = model.SensorNatureFromString(_type)
	}

	data, err := persist.NewSensorDataRepository().FindAllSensor(filter)
	if err != nil {
		fmt.Println(w, "Error retrieving data: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

	// TODO : MANAGE FUNCTION RETURN VALUE AND ERROR
}

//Example of getting data

func GetReadings(w http.ResponseWriter, r *http.Request) {
	sensorID := r.URL.Query().Get("sensorID")
	airportIATA := r.URL.Query().Get("airportIATA")
	_type := r.URL.Query().Get("type")

	filter := persist.Filter{
		From: time.Unix(0, 0),
		To:   time.Now(),
	}

	if sensorID != "" {
		filter.SensorId = sensorID
	}
	if airportIATA != "" {
		filter.AirportIATA = airportIATA
	}
	if _type != "" {
		filter.Type = model.SensorNatureFromString(_type)
	}

	data, err := persist.NewSensorDataRepository().FindAllReading(filter)
	if err != nil {
		fmt.Println(w, "Error retrieving data: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

	// TODO : MANAGE FUNCTION RETURN VALUE AND ERROR
}

func GlobalDailyAverage(w http.ResponseWriter, r *http.Request) {
	// TODO : implement
}

func DailyAverage(w http.ResponseWriter, r *http.Request) {
	// TODO : implement
}

func OnTimeList(w http.ResponseWriter, r *http.Request) {
	// TODO : implement
}
