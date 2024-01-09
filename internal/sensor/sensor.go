package sensor

import (
	"Airport_MQTT/internal/broker"
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/config/types"
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
	BrokerID    broker.Broker
	IataCode    string
	Measurement string
	Frequency   int
}

func NewSensor(sensorInterface SensorInterface, sensorId string, iataCode string, measurement string, frequency int) Sensor {
	cfg := config.LoadConfig(&types.ConfigFile{}, "config.yaml").(*types.ConfigFile)
	return Sensor{
		SensorID:        sensorId,
		BrokerID:        broker.NewBroker(cfg),
		IataCode:        iataCode,
		Measurement:     measurement,
		Frequency:       frequency,
		SensorInterface: sensorInterface,
	}
}

func (s Sensor) SendToBroker(data MeasurementData) {
	var valueStr string
	if data.Value != nil {
		valueStr = fmt.Sprintf("%v", *data.Value)
	} else {
		valueStr = "nil"
	}
	s.BrokerID.SendMessage(
		fmt.Sprintf("airport/%s/sensors%s/%s", s.IataCode, s.Measurement, s.SensorID),
		fmt.Sprintf("timestamp:%s\nvalue:%d\n", data.Timestamp, data.Value),
	)
	fmt.Printf("TypeMeasure: %s, Value: %s, Timestamp: %s\n", data.TypeMeasure, valueStr, data.Timestamp)
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
