package sensor

import (
	"Airport_MQTT/internal/config"
	"Airport_MQTT/internal/model"
	"Airport_MQTT/internal/mqttUtils"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/krisukox/google-flights-api/iata"
	"log/slog"
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
	DeviceId    string
	BrokerMqtt  mqtt.Client
	AirportIATA string
	Type        string
	Frequency   int
}

func NewSensor(sensorInterface SensorInterface, sensorId string, iataCode string, measurement string, frequency int) Sensor {
	return Sensor{
		DeviceId:        sensorId,
		BrokerMqtt:      mqttUtils.NewMqttClient(),
		AirportIATA:     iataCode,
		Type:            measurement,
		Frequency:       frequency,
		SensorInterface: sensorInterface,
	}
}

func (s Sensor) SendToBroker(data model.SensorData) {
	topic := mqttUtils.GetTopic(data, "sensors")
	message := fmt.Sprintf("{\"TypeMeasure\": \"%s\", \"Value\": %f, \"Timestamp\": \"%s\"}",
		data.Type, data.Value, data.Timestamp.Format("2006-01-02-15-04-05"))

	req := s.BrokerMqtt.Publish(
		topic, byte(config.GetMqttConfig().Client.QOS), false, message)
	req.Wait()
	fmt.Printf("TypeMeasure: %s, Value: %f, Timestamp: %s\n", data.Type, data.Value, data.Timestamp)
}

func (s Sensor) StartSensor() {
	location := iata.IATATimeZone(s.AirportIATA)
	if location.City == "" {
		slog.Error("IATA not supported : " + s.AirportIATA)
		return
	}

	ticker := time.NewTicker(time.Duration(s.Frequency) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.SendToBroker(s.getDataApi(s.Type, location.City))
		}
	}
}
