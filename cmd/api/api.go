package main

import (
	"Airport_MQTT/internal/api"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	http.Handle("/", router)

	router.HandleFunc("/sendData", api.SendDataExample)
	router.HandleFunc("/getSensor", api.GetSensor).
		Queries("sensorID", "{sensorID}", "airportIATA", "{airportIATA}", "type", "{type}")
	router.HandleFunc("/getReadings", api.GetReadings).
		Queries("sensorID", "{sensorID}", "airportIATA", "{airportIATA}", "type", "{type}")
	router.HandleFunc("/globalDailyAverage", api.GlobalDailyAverage)
	router.HandleFunc("/onTimeList", api.OnTimeList).Queries("type", "{type}")

	fmt.Printf("Serveur écoutant sur le port %d...\n", 8080)
	err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		return
	}
}
