package config

// Si possible de refactor pour ne pas avoir toutes les configs qui se répètent dans chaque executable
// Avec interface{} ??? (voir si possible) j'ai essayé, mais pas réussi

type Config struct {
	Datasource DatasourceConfig `yaml:"datasource"`
	Mqtt       MqttConfig       `yaml:"mqtt"`
	Sensor     SensorConfig     `yaml:"sensor"`
	API        APIConfig        `yaml:"api"`
	Web        WebConfig        `yaml:"web"`
	Alerts     []Alert          `yaml:"alerts"`
	File       File             `yaml:"file"`
}

type DatasourceConfig struct {
	Url      string `yaml:"url" env:"DATASOURCE_URL"`
	Username string `yaml:"username" env:"DATASOURCE_USERNAME"`
	Password string `yaml:"password" env:"DATASOURCE_PASSWORD"`
}

type MQTTBrokerConfig struct {
	Url      string `yaml:"url" env:"MQTT_BROKER_URL"`
	Username string `yaml:"username" env:"MQTT_BROKER_USERNAME"`
	Password string `yaml:"password" env:"MQTT_BROKER_PASSWORD"`
}

type MqttClientConfig struct {
	Id  string `yaml:"clientId" env:"MQTT_CLIENT_ID"`
	QOS int    `yaml:"qos" env:"MQTT_QOS"`
}

type MqttConfig struct {
	Broker MQTTBrokerConfig `yaml:"broker"`
	Client MqttClientConfig `yaml:"client"`
}

type SensorConfig struct {
	AirportIATA string `yaml:"airportIATA" env:"SENSOR_AIRPORT_IATA" validate:"required"`
	DeviceId    string `yaml:"deviceId" env:"SENSOR_DEVICE_ID"`
	SensorType  string `yaml:"sensorType" env:"SENSOR_TYPE"`
	Frequency   int    `yaml:"frequency" env:"SENSOR_FREQUENCY"`
}

type APIConfig struct {
	Url       string `yaml:"url" env:"API_URL"`
	SecretKey string `yaml:"secretKey" env:"API_SECRET_KEY"`
}

type WebConfig struct {
	Port int `yaml:"port"`
}

type Alert struct {
	Airport     string `yaml:"airport"`
	Temperature int    `yaml:"temperature"`
	Wind        int    `yaml:"wind"`
	Pressure    int    `yaml:"pressure"`
}

type File struct {
	Path string `yaml:"path"`
}
