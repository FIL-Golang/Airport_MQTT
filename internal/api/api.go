package api

import (
	"Airport_MQTT/internal/model"
	"Airport_MQTT/internal/persist"
	"encoding/json"
	"fmt"
	"log/slog"
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

// GetDistinctAirportCodes godoc
// @Summary Get distinct airport codes
// @Description Retrieve a list of all distinct airport IATA codes from the database.
// @ID get-distinct-airport-codes
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Failure 500 {object} ErrorResponse
// @Router /distinctAirportCodes [get]
func (controller *RestController) GetDistinctAirportCodes(w http.ResponseWriter, r *http.Request) {
	codes, err := controller.repository.GetDistinctAirportCodes()
	if err != nil {
		// Handle error
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("Error fetching airport codes: %v", err)})
		if err != nil {
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(codes)
	if err != nil {
		return
	}
}

// GetAllSensorsForAirport godoc
// @Summary Get all sensors for a specific airport
// @Description Retrieve all sensors associated with a given airport IATA code.
// @ID get-all-sensors-for-airport
// @Accept json
// @Produce json
// @Param airportIATA query string true "Airport IATA Code"
// @Success 200 {array} model.Sensor
// @Failure 500 {object} ErrorResponse
// @Router /sensorsForAirport [get]
func (controller *RestController) GetAllSensorsForAirport(w http.ResponseWriter, r *http.Request) {
	airportIATA := r.URL.Query().Get("airportIATA")

	filter := persist.Filter{AirportIATA: airportIATA}
	sensors, err := controller.repository.FindAllSensor(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to fetch sensors: " + err.Error()})
		if err != nil {
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode(sensors)
	if err != nil {
		return
	}
}

// GetAllReadingsForSensor godoc
// @Summary Get all readings for a specific sensor
// @Description Retrieve all readings from a specific sensor, optionally filtered by airport IATA code.
// @ID get-all-readings-for-sensor
// @Accept json
// @Produce json
// @Param sensorId query string true "Sensor ID"
// @Param airportIATA query string false "Airport IATA Code"
// @Success 200 {array} model.Sensor
// @Failure 500 {object} ErrorResponse
// @Router /readingsForSensor [get]
func (controller *RestController) GetAllReadingsForSensor(w http.ResponseWriter, r *http.Request) {
	sensorId := r.URL.Query().Get("sensorId")
	airportIATA := r.URL.Query().Get("airportIATA")

	filter := persist.Filter{
		SensorId:    sensorId,
		AirportIATA: airportIATA,
	}
	readings, err := controller.repository.FindAllReading(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to fetch readings: " + err.Error()})
		if err != nil {
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode(readings)
	if err != nil {
		return
	}
}

// GetLastReadingForSensor godoc
// @Summary Get the last reading for a specific sensor
// @Description Retrieve the last reading from a specific sensor, optionally filtered by airport IATA code.
// @ID get-last-reading-for-sensor
// @Accept json
// @Produce json
// @Param sensorId query string true "Sensor ID"
// @Param airportIATA query string false "Airport IATA Code"
// @Success 200 {object} model.SensorData
// @Failure 500 {object} ErrorResponse
// @Router /lastReadingForSensor [get]
func (controller *RestController) GetLastReadingForSensor(w http.ResponseWriter, r *http.Request) {
	sensorId := r.URL.Query().Get("sensorId")
	airportIATA := r.URL.Query().Get("airportIATA")

	filter := persist.Filter{
		SensorId:    sensorId,
		AirportIATA: airportIATA,
	}
	lastReading, err := controller.repository.GetLastReading(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to fetch the last reading: " + err.Error()})
		if err != nil {
			return
		}
		return
	}

	err = json.NewEncoder(w).Encode(lastReading)
	if err != nil {
		return
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
		err := json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			return
		}
		return
	}

	// Creation des objets time.Time pour le début et la fin de la journee
	debut := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	fin := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())

	reelType := model.SensorTypeFromString(typeParam)

	filter := persist.Filter{
		AirportIATA: airportIATA,
		Type:        reelType,
		From:        debut,
		To:          fin,
	}

	avg, err := controller.repository.GetAvg(filter)
	if err != nil {
		slog.Debug("Error while retrieving average: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{Error: "Problème lors de la récupération de la moyenne"}
		err := json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			return
		}
		return
	}

	response := SensorDataResponse{
		Jour: debut.Format("02/01/2006"),
		Avg:  avg,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

// OnTimeList godoc
// @Summary Get daily averages
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
		err := json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			return
		}
		return
	}

	fin, err := parseDate(paramTo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{Error: "Problème lors de la conversion de la date"}
		err := json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			return
		}
		return
	}

	reelType := model.SensorTypeFromString(typeParam)
	if reelType == model.Undefined {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{Error: "Problème lors de la conversion du type"}
		err := json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			return
		}
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
		err := json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}
