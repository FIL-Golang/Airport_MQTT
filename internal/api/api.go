package api

import (
	"Airport_MQTT/internal/model"
	"Airport_MQTT/internal/persist"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func parseDate(dateStr string) (time.Time, error) {
	formatDate := "02-01-2006"
	return time.Parse(formatDate, dateStr)
}

type SensorDataResponse struct {
	Jour string          `json:"jour"`
	Avg  []model.Average `json:"avg"`
}

type GlobalSensorDataResponse struct {
	Jour        string          `json:"jour"`
	AvgTemp     []model.Average `json:"avgTemperature"`
	AvgPressure []model.Average `json:"avgPressure"`
	AvgWind     []model.Average `json:"avgWind"`
}

type ListDataResponse struct {
	Jour string         `json:"jour"`
	Avg  []model.Sensor `json:"avg"`
}

func DailyAverage(w http.ResponseWriter, r *http.Request) {
	typeParam := r.URL.Query().Get("type")
	paramDay := r.URL.Query().Get("day")

	// Convertion de la chaîne de caractères de la date en objet time.Time
	date, err := parseDate(paramDay)
	if err != nil {
		fmt.Println("Erreur lors de la conversion de la date :", err)
		return
	}

	// Creation des objets time.Time pour le début et la fin de la journee
	debut := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	fin := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())

	repo := persist.NewSensorDataRepository()

	if typeParam == "" {
		temperature := model.SensorNatureFromString("temperature")
		filterTemp := persist.Filter{
			Type: temperature,
			From: debut,
			To:   fin,
		}
		avgTemp, err := repo.GetAvg(filterTemp)
		if err != nil {
			fmt.Println("Erreur lors de la récupération de la moyenne de rempérature :", err)
			return
		}

		pressure := model.SensorNatureFromString("pressure")
		filterPress := persist.Filter{
			Type: pressure,
			From: debut,
			To:   fin,
		}
		avgPress, err := repo.GetAvg(filterPress)

		windSpeed := model.SensorNatureFromString("wind_speed")
		filterWind := persist.Filter{
			Type: windSpeed,
			From: debut,
			To:   fin,
		}
		avgWind, err := repo.GetAvg(filterWind)

		response := GlobalSensorDataResponse{
			Jour:        debut.Format("02/01/2006"),
			AvgTemp:     avgTemp,
			AvgPressure: avgPress,
			AvgWind:     avgWind,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		//Creation du type
		reelType := model.SensorNatureFromString(typeParam)
		if reelType == model.Undefined {
			fmt.Println("Erreur lors de la conversion du type :", err)
			return
		}

		filter := persist.Filter{
			Type: reelType,
			From: debut,
			To:   fin,
		}

		avg, err := repo.GetAvg(filter)
		if err != nil {
			fmt.Println("Erreur lors de la récupération de la moyenne :", err)
			return
		}

		response := SensorDataResponse{
			Jour: debut.Format("02/01/2006"),
			Avg:  avg,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func OnTimeList(w http.ResponseWriter, r *http.Request) {
	typeParam := r.URL.Query().Get("type")
	paramDay := r.URL.Query().Get("day")

	// Convertion de la chaîne de caractères de la date en objet time.Time
	date, err := parseDate(paramDay)
	if err != nil {
		fmt.Println("Erreur lors de la conversion de la date :", err)
		return
	}

	// Creation des objets time.Time pour le début et la fin de la journee
	debut := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	fin := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())

	reelType := model.SensorNatureFromString(typeParam)
	if reelType == model.Undefined {
		fmt.Println("Erreur lors de la conversion du type :", err)
		return
	}
	filter := persist.Filter{
		Type: reelType,
		From: debut,
		To:   fin,
	}

	data, err := persist.NewSensorDataRepository().FindAllReading(filter)
	if err != nil {
		fmt.Println(w, "Error retrieving data: %v", err)
		return
	}

	response := ListDataResponse{
		Jour: debut.Format("02/01/2006"),
		Avg:  data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
