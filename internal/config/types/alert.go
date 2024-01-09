package types

type AlertConfigFile struct {
	Alerts []Alert `yaml:"alerts"`
}

type Alert struct {
	Airport     string `yaml:"airport"`
	Temperature int    `yaml:"temperature"`
	Wind        int    `yaml:"wind"`
	Pressure    int    `yaml:"pressure"`
}
