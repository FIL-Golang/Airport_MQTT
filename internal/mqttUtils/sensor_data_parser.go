package mqttUtils

import (
	"Airport_MQTT/internal/model"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"
)

type payload struct {
	Value     float32 `json:"value"`
	Timestamp string  `json:"timestamp"`
}

const (
	posAirportIATA = 2
	posSensorId    = 5
	posSensorType  = 4
)

func Parse(message mqtt.Message) (error, model.SensorData) {
	sensorData := model.SensorData{}
	err := parseTopic(message, &sensorData)
	if err != nil {
		return err, sensorData
	}
	err = parsePayload(message, &sensorData)
	if err != nil {
		return err, sensorData
	}
	return nil, sensorData
}
func parseTopic(message mqtt.Message, data *model.SensorData) error {
	splitTopic := strings.Split(message.Topic(), "/")
	data.SensorId = splitTopic[posSensorId]
	data.Nature = model.SensorNatureFromString(splitTopic[posSensorType])
	data.AirportIATA = splitTopic[posAirportIATA]
	return nil
}

func parsePayload(message mqtt.Message, data *model.SensorData) error {
	pay := payload{}
	err := json.Unmarshal(message.Payload(), &pay)
	if err != nil {
		return err
	}
	if pay.Value == 0 || pay.Timestamp == "" {
		return nil
	}
	data.Value = pay.Value
	return nil
}

func GetTopic(sensorData model.SensorData, typ string) string {
	return "/airports/" + sensorData.AirportIATA + "/" + typ + "/" + sensorData.Nature.String() + "/" + sensorData.SensorId
}

func GetSensorsTopic(sensor model.SensorData) string {
	return GetTopic(sensor, "sensors")
}

func GetAlertsTopic(sensor model.SensorData) string {
	return GetTopic(sensor, "alerts")
}

func GetPayload(reading model.SensorData) []byte {
	pay := payload{
		Value:     reading.Value,
		Timestamp: reading.Timestamp.Format("2006-01-02-15-04-05"),
	}
	bytePayload, _ := json.Marshal(pay)
	return bytePayload
}

func convertStringToTime(timestamp string) (time.Time, error) {
	return time.Parse("2006-01-02-15-04-05", timestamp)
}
