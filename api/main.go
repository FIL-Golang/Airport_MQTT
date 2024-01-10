package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func test(w http.ResponseWriter, r *http.Request) {
	test := Person{ID: 1, Name: "allo"}
	JsonData, err := json.Marshal(test)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Configuration de l'en-tête Content-Type pour indiquer que la réponse est au format JSON
	w.Header().Set("Content-Type", "application/json")

	//renvoi des données
	w.Write(JsonData)
}

func globalDailyAverage(w http.ResponseWriter, r *http.Request) {

}

func dailyAverage(w http.ResponseWriter, r *http.Request) {

}

func onTimeList(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/globalDailyAverage", globalDailyAverage)
	r.HandleFunc("/dailyAverage/{type}", dailyAverage)
	r.HandleFunc("/onTimeList/{type}", onTimeList)
	http.Handle("/", r)

	// Démarrer le serveur sur le port 8080
	port := 8080
	fmt.Printf("Serveur écoutant sur le port %d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
