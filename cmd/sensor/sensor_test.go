package main

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "-config", "test_config.yaml"}

	config := loadConfig()

	if config.SensorId == "" || config.IataCode == "" || config.Measurement == "" {
		t.Errorf("Config fields are not populated correctly")
	}
}

func TestValidateConfig(t *testing.T) {
	validUUID := "7f9c4949-1065-4ca1-b040-991370797d8f"
	testConfig := Config{
		SensorId:    validUUID,
		IataCode:    "ABC",
		Measurement: "temp",
		Frequency:   5,
	}

	err := validateConfig(&testConfig)
	if err != nil {
		t.Errorf("validateConfig failed for valid config: %v, error: %v", testConfig, err)
	}
}

func TestInvalidateConfig(t *testing.T) {
	if os.Getenv("GO_TEST_SUBPROCESS") == "1" {
		validUUID := "7f9c4949-1065-4ca1-b040-991370797d8f"
		testConfig := Config{
			SensorId:    validUUID,
			IataCode:    "",
			Measurement: "temp",
			Frequency:   5,
		}
		validateConfig(&testConfig)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestInvalidateConfig")
	cmd.Env = append(os.Environ(), "GO_TEST_SUBPROCESS=1")
	err := cmd.Run()

	var e *exec.ExitError
	if errors.As(err, &e) && !e.Success() {
		return
	}

	t.Fatalf("validateConfig did not call os.Exit(1) as expected; err: %v", err)
}
