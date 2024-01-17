package main

import (
	"Airport_MQTT/cmd/api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/globalDailyAverage", api.GlobalDailyAverage)
	r.HandleFunc("/dailyAverage", api.DailyAverage).Queries("from", "{from}", "to", "{to}", "type", "{type}")
	r.HandleFunc("/onTimeList", api.OnTimeList).Queries("from", "{from}", "to", "{to}", "type", "{type}")

	r.HandleFunc("/sendData", api.SendDataExample)
	r.HandleFunc("/getSensor", api.GetSensor).
		Queries("sensorID", "{sensorID}", "airportIATA", "{airportIATA}", "type", "{type}")
	r.HandleFunc("/getReadings", api.GetReadings).
		Queries("sensorID", "{sensorID}", "airportIATA", "{airportIATA}", "type", "{type}")
	r.HandleFunc("/globalDailyAverage", api.GlobalDailyAverage)
	r.HandleFunc("/onTimeList", api.OnTimeList).Queries("type", "{type}")

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