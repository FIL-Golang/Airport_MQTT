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
	r.HandleFunc("/processJSON", api.ProcessJSONData).Methods("POST") // Nouvelle route pour traiter le JSON dans le corps
	http.Handle("/", r)

	// Démarrer le serveur sur le port 8080
	port := 8080
	fmt.Printf("Serveur écoutant sur le port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
