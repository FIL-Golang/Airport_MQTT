package main

import (
	"github.com/google/uuid"
	"testing"
)

func TestParseArgsWithValidUUID(t *testing.T) {
	validUUID := uuid.New().String()
	args := Args{
		SensorId:    validUUID,
		IataCode:    "XYZ",
		Measurement: "Temperature",
		Frequency:   "5",
	}

	sensorId, iataCode, _, _, err := parseArgs(args)
	if err != nil {
		t.Fatalf("parseArgs returned an error: %v", err)
	}

	if sensorId != validUUID {
		t.Errorf("Expected sensorId to be '%s', got %s", validUUID, sensorId)
	}
	if iataCode != "XYZ" {
		t.Errorf("Expected iataCode to be 'XYZ', got %s", iataCode)
	}
}

func TestParseArgsWithInvalidUUID(t *testing.T) {
	args := Args{
		SensorId:    "1234",
		IataCode:    "XYZ",
		Measurement: "Temperature",
		Frequency:   "5",
	}

	sensorId, _, _, _, err := parseArgs(args)
	if err != nil {
		t.Fatalf("parseArgs returned an error: %v", err)
	}

	if _, uuidErr := uuid.Parse(sensorId); uuidErr != nil {
		t.Errorf("Expected sensorId to be a valid UUID, got %s", sensorId)
	}
}
