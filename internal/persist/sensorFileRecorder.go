package persist

import (
	"Airport_MQTT/internal/model"
)

type SensorFileRecorder interface {
	Store(data model.SensorData) (savedData model.SensorData, err error)
}

type sensorFileRecorder struct {
	directory string
}

func NewSensorFileRecorder() SensorFileRecorder {
	return &sensorFileRecorder{
		directory: "",
	}
}

func (r *sensorFileRecorder) Store(data model.SensorData) (savedData model.SensorData, err error) {
	//_, err = r.collection.InsertOne(r.ctx, data)
	if err != nil {
		return data, err
	}
	return data, nil
}
