package main

import (
	"Airport_MQTT/internal/api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/dailyAverage", api.DailyAverage).
		Queries("day", "{day}")
	r.HandleFunc("/onTimeList", api.OnTimeList).
		Queries("day", "{day}", "type", "{type}")

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
