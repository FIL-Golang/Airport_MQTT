package main

import (
	_ "Airport_MQTT/cmd/web_api/docs"
	"Airport_MQTT/internal/api"
	"Airport_MQTT/internal/config"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	controller := api.NewRestController()

	r.HandleFunc("/dailyAverage", controller.DailyAverage).
		Queries("day", "{day}").
		Methods("GET")

	r.HandleFunc("/onTimeList", controller.OnTimeList).
		Queries("from", "{from}", "to", "{to}", "type", "{type}").
		Methods("GET")

	r.HandleFunc("/distinctAirportCodes", controller.GetDistinctAirportCodes)

	r.HandleFunc("/sensorsForAirport", controller.GetAllSensorsForAirport).
		Queries("airportIATA", "{airportIATA}")

	r.HandleFunc("/readingsForSensor", controller.GetAllReadingsForSensor).
		Queries("sensorId", "{sensorId}", "airportIATA", "{airportIATA}")

	//Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.Handle("/", r)

	// Démarrer le serveur sur le port 8080
	port := config.GetWebConfig().Port
	slog.Info("Starting server on port " + fmt.Sprintf("%d", port))
	//http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		slog.Debug("Error starting server: " + err.Error())
		return
	}
}
