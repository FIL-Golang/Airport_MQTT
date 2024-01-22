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
	err := parseTopic(message.Topic(), &sensorData)
	if err != nil {
		return err, sensorData
	}
	err = parsePayload(message.Payload(), &sensorData)
	if err != nil {
		return err, sensorData
	}
	return nil, sensorData
}
func parseTopic(topic string, data *model.SensorData) error {
	splitTopic := strings.Split(topic, "/")
	data.SensorId = splitTopic[posSensorId]
	data.Type = model.SensorTypeFromString(splitTopic[posSensorType])
	data.AirportIATA = splitTopic[posAirportIATA]
	return nil
}

func parsePayload(payloadByte []byte, data *model.SensorData) error {
	pay := payload{}
	err := json.Unmarshal(payloadByte, &pay)
	if err != nil {
		return err
	}
	if pay.Value == 0 || pay.Timestamp == "" {
		return nil
	}
	data.Value = pay.Value
	data.Timestamp, err = convertStringToTime(pay.Timestamp)
	if err != nil {
		return err
	}

	return nil
}

func GetTopic(sensorData model.SensorData, typ string) string {
	return "/airports/" + sensorData.AirportIATA + "/" + typ + "/" + sensorData.Type.String() + "/" + sensorData.SensorId
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
