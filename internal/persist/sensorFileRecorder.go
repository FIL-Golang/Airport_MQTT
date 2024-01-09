package persist

import (
	"Airport_MQTT/internal/model"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type SensorFileRecorder interface {
	Store(data model.SensorData) (err error)
}

type sensorFileRecorder struct {
	directory string
}

func NewSensorFileRecorder() SensorFileRecorder {
	return &sensorFileRecorder{
		directory: "",
	}
}

func (r *sensorFileRecorder) Store(data model.SensorData) (err error) {
	err = writeSensorData(data)
	if err != nil {
		return err
	}
	return nil
}

func writeSensorData(data model.SensorData) error {
	// Construire le chemin du dossier records pour l'aéroport
	airportDir := filepath.Join("records", data.AirportIATA)
	if err := os.MkdirAll(airportDir, os.ModePerm); err != nil {
		return err
	}

	// Construire le chemin du dossier pour le type de donnée
	dataTypeDir := filepath.Join(airportDir, data.Nature.String())
	if err := os.MkdirAll(dataTypeDir, os.ModePerm); err != nil {
		return err
	}

	// Construire le chemin complet du fichier CSV
	fileName := filepath.Join(dataTypeDir, fmt.Sprintf("%s.csv", data.Timestamp))

	// Ouvrir le fichier CSV en mode append ou créer s'il n'existe pas encore
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Créer un nouveau writer pour le fichier CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Si le fichier est nouvellement créé, écrire l'en-tête du fichier CSV
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if fileInfo.Size() == 0 {
		header := []string{"Timestamp", "Value"}
		writer.Write(header)
	}

	// Écrire les données dans le fichier CSV
	row := []string{data.Timestamp, fmt.Sprintf("%.2f", data.Value)}
	writer.Write(row)

	return nil
}
