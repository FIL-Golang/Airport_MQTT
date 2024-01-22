package config

import (
	"log/slog"
	"os"
	"reflect"

	"gopkg.in/yaml.v3"
)

var config Config

func LoadConfig() {
	args := os.Args
	if len(args) != 2 {
		slog.Warn("Usage: <exe> <config file path>")
		os.Exit(1)
	}
	config = Config{}

	err := LoadValuesFromYaml(&config, args[1])
	if err != nil {
		panic(err)
	}
	err = LoadValuesFromEnv(&config)
	if err != nil {
		panic(err)
	}
}

// LoadValuesFromEnv LoadEnvValue looks for fields with the tag env and replace the value by the value of the environment variable
// with the name specified in the tag if it exists. If the environment variable doesn't exist, the value is not changed
// it is recursive and will explore all the fields of the struct and sub structs
func LoadValuesFromEnv(configPtr interface{}) error {
	configType := reflect.TypeOf(configPtr).Elem()
	configValue := reflect.ValueOf(configPtr).Elem()
	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		fieldValue := configValue.Field(i)
		if field.Type.Kind() == reflect.Struct {
			err := LoadValuesFromEnv(fieldValue.Addr().Interface().(interface{}))
			if err != nil {
				return err
			}
		} else {
			tag := field.Tag.Get("env")
			if tag != "" {
				envValue := os.Getenv(tag)
				if envValue != "" {
					fieldValue.SetString(envValue)
				}
			}
		}
	}

	return nil
}

func LoadValuesFromYaml(configPtr *Config, path string) error {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, configPtr)
	if err != nil {
		return err
	}

	return nil
}

func GetDatasourceConfig() DatasourceConfig {
	return config.Datasource
}

func GetMqttConfig() MqttConfig {
	return config.Mqtt
}

func GetAlerts() []Alert { return config.Alerts }

func GetSensorConfig() SensorConfig {
	return config.Sensor
}

func GetApiConfig() APIConfig {
	return config.API
}

func GetFileConfig() File {
	return config.File
}

func GetWebConfig() WebConfig {
	return config.Web
}
