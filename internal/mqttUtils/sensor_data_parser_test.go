package mqttUtils

import (
	"Airport_MQTT/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseTopic(t *testing.T) {
	//prepare
	topic := "/airports/NTE/sensors/temperature/103F6526-C630-4D56-9523-160BE051A956"
	sensorData := model.SensorData{}

	//act
	err := parseTopic(topic, &sensorData)
	if err != nil {
		t.Errorf("Error parsing topic: %v", err)
	}

	//assert
	assert.Equal(t, "103F6526-C630-4D56-9523-160BE051A956", sensorData.SensorId)
	assert.Equal(t, "NTE", sensorData.AirportIATA)
	assert.Equal(t, model.Temperature, sensorData.Type)
}

func TestParsePayload(t *testing.T) {
	//prepare
	payload := []byte(`{"value":12.3,"timestamp":"2024-01-15-22-30-40"}`)
	sensorData := model.SensorData{}

	//act
	err := parsePayload(payload, &sensorData)
	if err != nil {
		t.Errorf("Error parsing payload: %v", err)
	}

	//assert
	assert.Equal(t, float32(12.3), sensorData.Value)
	assert.Equal(t, "2024-01-15 22:30:40 +0000 UTC", sensorData.Timestamp.String())
}

func TestGetAlertsTopic(t *testing.T) {
	//prepare
	sensorData := model.SensorData{
		SensorId:    "103F6526-C630-4D56-9523-160BE051A956",
		AirportIATA: "NTE",
		Type:        model.Temperature,
		Value:       12.3,
		Timestamp:   time.Date(2024, 1, 15, 22, 30, 40, 0, time.UTC),
	}

	//act
	topic := GetAlertsTopic(sensorData)

	//assert
	assert.Equal(t, "/airports/NTE/alerts/temperature/103F6526-C630-4D56-9523-160BE051A956", topic)
}

func TestGetSensorsTopic(t *testing.T) {
	//prepare
	sensorData := model.SensorData{
		SensorId:    "103F6526-C630-4D56-9523-160BE051A956",
		AirportIATA: "NTE",
		Type:        model.Temperature,
		Value:       12.3,
		Timestamp:   time.Date(2024, 1, 15, 22, 30, 40, 0, time.UTC),
	}

	//act
	topic := GetSensorsTopic(sensorData)

	//assert
	assert.Equal(t, "/airports/NTE/sensors/temperature/103F6526-C630-4D56-9523-160BE051A956", topic)
}

func TestGetPayload(t *testing.T) {
	//prepare
	sensorData := model.SensorData{
		SensorId:    "103F6526-C630-4D56-9523-160BE051A956",
		AirportIATA: "NTE",
		Type:        model.Temperature,
		Value:       12.3,
		Timestamp:   time.Date(2024, 1, 15, 22, 30, 40, 0, time.UTC),
	}

	//act
	payload := GetPayload(sensorData)

	//assert
	assert.Equal(t, []byte(`{"value":12.3,"timestamp":"2024-01-15-22-30-40"}`), payload)
}
