package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func parseDate(dateStr string) (time.Time, error) {
	formatDate := "02-01-2006" // Mettez ici votre format de date attendu
	return time.Parse(formatDate, dateStr)
}

func GlobalDailyAverage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bonjour depuis GlobalDailyAverage!")
	//persist.NewSensorDataRepository().GetAvg()
}

func DailyAverage(w http.ResponseWriter, r *http.Request) {
	typeParam := r.URL.Query().Get("type")
	parametreDebut := r.URL.Query().Get("from")
	parametreFin := r.URL.Query().Get("to")

	// Conversion des paramètres en objets time.Time en utilisant la fonction
	debut, err := parseDate(parametreDebut)
	if err != nil {
		fmt.Println("Erreur lors de la conversion de la date de début :", err)
		return
	}

	fin, err := parseDate(parametreFin)
	if err != nil {
		fmt.Println("Erreur lors de la conversion de la date de fin :", err)
		return
	}
	fmt.Fprint(w, debut, fin, typeParam)
	//persist.NewSensorDataRepository().GetAvg()
}

func OnTimeList(w http.ResponseWriter, r *http.Request) {
	typeParam := r.URL.Query().Get("type")
	//persist.NewSensorDataRepository().FindAllReading()
	fmt.Fprint(w, typeParam)
}

type JsonRequest struct {
	Type string `json:"type"`
	// Ajoutez d'autres champs selon votre besoin
}

func ProcessJSONData(w http.ResponseWriter, r *http.Request) {
	// Vérifier que la méthode est POST
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Décodez le JSON brut du corps de la requête
	var requestData JsonRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
		return
	}

	// Vous pouvez maintenant utiliser les données décodées (requestData)
	fmt.Fprintf(w, "Données JSON reçues: %+v", requestData)
}
