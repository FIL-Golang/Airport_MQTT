package sensor

import (
	"fmt"
	"github.com/krisukox/google-flights-api/iata"
	"time"
)

type MeasurementData struct {
	TypeMeasure string
	Value       *float64
	Timestamp   string
}

type SensorInterface interface {
	StartSensor()
	SendToBroker(measureData MeasurementData)
	getDataApi(measure string, city string) MeasurementData
}

type Sensor struct {
	SensorInterface
	SensorID    string
	BrokerID    string
	IataCode    string
	Measurement string
	Frequency   int
}

func NewSensor(sensorInterface SensorInterface, sensorId string, iataCode string, measurement string, frequency int) Sensor {
	return Sensor{
		SensorID:        sensorId,
		BrokerID:        "1", //NewBrokerId(),
		IataCode:        iataCode,
		Measurement:     measurement,
		Frequency:       frequency,
		SensorInterface: sensorInterface,
	}
}

func (s Sensor) SendToBroker(data MeasurementData) {
	//s.BrokerID.SendMessage()
	fmt.Println(data)
}

func (s Sensor) StartSensor() {
	location := iata.IATATimeZone(s.IataCode)
	if location.City == "" {
		fmt.Println("IATA not supported : ", s.IataCode)
		return
	}

	ticker := time.NewTicker(time.Duration(s.Frequency) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.SendToBroker(s.getDataApi(s.Measurement, location.City))
		}
	}
}
