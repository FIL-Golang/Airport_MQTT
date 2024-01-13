package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var testDataDir = "../../test/data"

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestLoadValueFromEnv(t *testing.T) {
	os.Setenv("DATASOURCE_URL", "http://localhost:8080/DBTest")
	os.Setenv("DATASOURCE_USERNAME", "admin")
	os.Setenv("DATASOURCE_PASSWORD", "admin")

	os.Setenv("MQTT_BROKER_PASSWORD", "adminMqtt")

	confStruct := Config{}

	err := LoadValuesFromEnv(&confStruct)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(confStruct)

	assert.Equal(t, "http://localhost:8080/DBTest", confStruct.Datasource.Url)
	assert.Equal(t, "admin", confStruct.Datasource.Username)
	assert.Equal(t, "admin", confStruct.Datasource.Password)

	assert.Equal(t, "adminMqtt", confStruct.Mqtt.Broker.Password)
}

func TestLoadValueFromYaml(t *testing.T) {

	confStruct := Config{}

	err := LoadValuesFromYaml(&confStruct, testDataDir+"/config.yaml")
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	assert.Equal(t, "http://localhost:8080/DBTest", confStruct.Datasource.Url)
	assert.Equal(t, "user", confStruct.Datasource.Username)
	assert.Equal(t, "", confStruct.Datasource.Password)

	assert.Equal(t, "tcp://localhost:1883", confStruct.Mqtt.Broker.Url)
}

func TestLoadConfig(t *testing.T) {
	os.Setenv("DATASOURCE_URL", "http://localhost:8080/DBTest")
	os.Setenv("DATASOURCE_USERNAME", "admin")
	os.Setenv("DATASOURCE_PASSWORD", "admin")

	LoadConfig(testDataDir + "/config.yaml")

	assert.Equal(t, "http://localhost:8080/DBTest", config.Datasource.Url)
	assert.Equal(t, "admin", config.Datasource.Username)
	assert.Equal(t, "admin", config.Datasource.Password)
}

func TestShouldGivePriorityToEnv(t *testing.T) {
	os.Setenv("DATASOURCE_USERNAME", "")         //should let the yaml value
	os.Setenv("DATASOURCE_PASSWORD", "password") //should override the yaml value

	LoadConfig(testDataDir + "/config.yaml")

	assert.Equal(t, "http://localhost:8080/DBTest", config.Datasource.Url)
	assert.Equal(t, "user", config.Datasource.Username)
	assert.Equal(t, "password", config.Datasource.Password)
}
