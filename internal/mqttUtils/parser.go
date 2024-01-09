package mqttUtils

import (
	"Airport_MQTT/internal/model"
	"Airport_MQTT/internal/utils"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
)

type Parser interface {
	Parse(message mqtt.Message) model.SensorData
}

type parser struct {
	sensorId    string
	sensorType  int
	airportIATA string
	sensorValue float32
}

func NewParser() Parser {
	return &parser{}
}

func (p *parser) parseTopic(message mqtt.Message) {
	fmt.Println(message.Topic())

	splitTopic := strings.Split(message.Topic(), "/")
	p.sensorId = splitTopic[5]
	p.sensorType = model.SensorNatureFromString(splitTopic[4])
	p.airportIATA = splitTopic[2]
}

func (p *parser) parsePayload(message mqtt.Message) {
	p.sensorValue = utils.ByteToFloat32(message.Payload())
}

func (p *parser) Parse(message mqtt.Message) model.SensorData {
	p.parseTopic(message)
	p.parsePayload(message)
	return model.SensorData{
		SensorId: p.sensorId,
		Nature:   p.sensorType,
		CodeIATA: p.airportIATA,
		Value:    p.sensorValue,
	}
}
