package main

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/model"
	"Airport_MQTT/internal/persist"
	"github.com/google/uuid"
	"log/slog"
	"math/rand"
	"time"
)

func init() {
	config.LoadConfig()
}

func main() {
	repo := persist.NewSensorDataRepository()
	from := time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC)
	to := time.Date(2024, 01, 25, 0, 0, 0, 0, time.UTC)
	sensorData := genSensorDataFromTo(from, to)
	for _, data := range sensorData {
		if err := repo.Store(data); err != nil {
			slog.Error("Error while storing sensor data: " + err.Error())
		}
	}
	slog.Info("Sensor data stored")
}

func genSensorDataFromTo(from time.Time, to time.Time) []model.SensorData {
	var sensorData []model.SensorData

	sensorMap := []map[string]string{
		{
			"sensorId": uuid.New().String(),
			"type":     model.Temperature.String(),
			"iata":     "NTE",
		},
		{
			"sensorId": uuid.New().String(),
			"type":     model.Pressure.String(),
			"iata":     "NTE",
		},
		{
			"sensorId": uuid.New().String(),
			"type":     model.WindSpeed.String(),
			"iata":     "NTE",
		},
	}

	for from.Before(to) {
		from = from.Add(time.Minute * 60)

		for _, sensor := range sensorMap {
			sensorData = append(sensorData, genSensorData(from, model.SensorTypeFromString(sensor["type"]), sensor["iata"], sensor["sensorId"]))
		}

	}
	return sensorData
}

func genSensorData(date time.Time, ty model.Type, iata string, sensorId string) model.SensorData {
	var sensorData model.SensorData
	var value float32
	switch ty {
	case model.Temperature:
		value = getRandomTemperature()
	case model.Pressure:
		value = getRandomPressure()
	case model.WindSpeed:
		value = getRandomWindSpeed()
	}
	sensorData = model.SensorData{
		Type:        ty,
		Value:       value,
		Timestamp:   date,
		AirportIATA: iata,
		SensorId:    sensorId,
	}
	return sensorData
}

func getRandomTemperature() float32 {
	return rand.Float32()*6 + 8
}

func getRandomPressure() float32 {
	return rand.Float32()*50 + 950
}

func getRandomWindSpeed() float32 {
	return rand.Float32()*10 + 5
}

func getRandomAirportIATA() string {
	return [...]string{
		"CDG",
		"ORY",
		"CPH",
		"NTE",
	}[rand.Intn(4)]
}
