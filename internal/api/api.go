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

type ErrorResponse struct {
	Error string `json:"error"`
}

type RestController struct {
	repository persist.SensorDataRepository
}

func NewRestController() *RestController {
	repository := persist.NewSensorDataRepository()
	return &RestController{
		repository: repository,
	}
}

// DailyAverage godoc
// @Summary Get daily averages
// @Description Get daily averages for temperature, pressure, and wind speed or everything.
// @ID get-daily-averages
// @Accept json
// @Produce json
// @Param day query string true "Date in the format '02-01-2006'"
// @Param type query string false "Type of sensor data (temperature, pressure, wind_speed)"
// @Success 200 {object} GlobalSensorDataResponse
// @Failure 400 {object} ErrorResponse
// @Router /dailyAverage [get]
func (controller *RestController) DailyAverage(w http.ResponseWriter, r *http.Request) {
	typeParam := r.URL.Query().Get("type")
	paramDay := r.URL.Query().Get("day")
	airportIATA := r.URL.Query().Get("airportIATA")

	// Convertion de la chaîne de caractères de la date en objet time.Time
	date, err := parseDate(paramDay)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{Error: "Problème de conversion de la date"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Creation des objets time.Time pour le début et la fin de la journee
	debut := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	fin := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())

	reelType := model.SensorNatureFromString(typeParam)

	filter := persist.Filter{
		AirportIATA: airportIATA,
		Type:        reelType,
		From:        debut,
		To:          fin,
	}

	avg, err := controller.repository.GetAvg(filter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		errorResponse := ErrorResponse{Error: "Problème lors de la récupération de la moyenne"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := SensorDataResponse{
		Jour: debut.Format("02/01/2006"),
		Avg:  avg,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// OnTimeList godoc
// @Summary Get readings by type
// @Description Get every measures by a type.
// @ID on-time-list
// @Accept json
// @Produce json
// @Param day query string true "Date in the format '02-01-2006'"
// @Param type query string true "Type of sensor data (temperature, pressure, wind_speed)"
// @Success 200 {object} GlobalSensorDataResponse
// @Failure 400 {object} ErrorResponse
// @Router /onTimeList [get]
func (controller *RestController) OnTimeList(w http.ResponseWriter, r *http.Request) {
	typeParam := r.URL.Query().Get("type")
	paramFrom := r.URL.Query().Get("from")
	paramTo := r.URL.Query().Get("to")

	// Convertion de la chaîne de caractères de la date en objet time.Time
	debut, err := parseDate(paramFrom)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{Error: "Problème lors de la conversion de la date"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	fin, err := parseDate(paramTo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{Error: "Problème lors de la conversion de la date"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	reelType := model.SensorNatureFromString(typeParam)
	if reelType == model.Undefined {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{Error: "Problème lors de la conversion du type"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	filter := persist.Filter{
		Type: reelType,
		From: debut,
		To:   fin,
	}

	data, err := controller.repository.FindAllReading(filter)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{Error: "Problème lors de la récupération des mesures"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
