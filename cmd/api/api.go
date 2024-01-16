package api

import (
	"Airport_MQTT/internal/persist"
	"encoding/json"
	"fmt"
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

func GlobalDailyAverage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bonjour depuis GlobalDailyAverage!")
}

func DailyAverage(w http.ResponseWriter, r *http.Request) {
	typeParam := r.URL.Query().Get("type")
	fmt.Fprint(w, typeParam)
	persist.NewSensorDataRepository().GetAvg()
}

func OnTimeList(w http.ResponseWriter, r *http.Request) {
	typeParam := r.URL.Query().Get("type")
	fmt.Fprint(w, typeParam)
}
