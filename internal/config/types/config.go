package types

type ConfigFile struct {
	Name string `yaml:"name"`

	Datasource struct {
		Url      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"datasource"`

	Mqtt struct {
		Url      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Topic    string `yaml:"topic"`
	} `yaml:"mqttBroker"`

	Web struct {
		Port int `yaml:"port"`
	} `yaml:"web"`
}
