package main

import (
	"Airport_MQTT/internal/api"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	controller := api.NewRestController()
	r.HandleFunc("/dailyAverage", controller.DailyAverage).
		Queries("day", "{day}")
	r.HandleFunc("/onTimeList", controller.OnTimeList).
		Queries("from", "{from}", "to", "{to}", "type", "{type}")
	r.HandleFunc("/distinctAirportCodes", controller.GetDistinctAirportCodes)
	r.HandleFunc("/sensorsForAirport", controller.GetAllSensorsForAirport).
		Queries("airportIATA", "{airportIATA}")
	r.HandleFunc("/readingsForSensor", controller.GetAllReadingsForSensor).
		Queries("sensorId", "{sensorId}", "airportIATA", "{airportIATA}")

	//Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.Handle("/", r)

	// Démarrer le serveur sur le port 8080
	port := 8080
	fmt.Printf("Serveur écoutant sur le port %d...\n", port)
	//http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}
}
