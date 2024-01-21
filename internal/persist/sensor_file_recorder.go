package persist

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/model"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
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
	fileConfig := config.GetFileConfig()
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filepath.Dir(filepath.Dir(b)))
	err := godotenv.Load(filepath.Join(basepath, "../../.env"))

	airportDir := filepath.Join(filepath.Join(basepath, fileConfig.Path), data.AirportIATA)
	if err := os.MkdirAll(airportDir, os.ModePerm); err != nil {
		return err
	}

	dataTypeDir := filepath.Join(airportDir, data.Nature.String())
	if err := os.MkdirAll(dataTypeDir, os.ModePerm); err != nil {
		return err
	}

	fileName := filepath.Join(dataTypeDir, fmt.Sprintf("%s-%s-%s.csv", data.AirportIATA, data.Timestamp.Format("2006-01-02"), data.Nature.String()))
	fmt.Println(fileName)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if fileInfo.Size() == 0 {
		header := []string{"Timestamp", "Value", "SensorId"}
		writer.Write(header)
	}

	row := []string{data.Timestamp.Format("15:04:05"), fmt.Sprintf("%.2f", data.Value), data.SensorId}
	writer.Write(row)

	return nil
}
