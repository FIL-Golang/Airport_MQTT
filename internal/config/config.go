package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"regexp"
)

type Config struct {
	Datasource struct {
		Url      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"datasource"`

	Web struct {
		Port int `yaml:"port"`
	} `yaml:"web"`
}

func LoadConfig() (error, Config) {
	yamlFile, err := os.ReadFile("config/config.yaml")
	cfg := Config{}

	re := regexp.MustCompile(`\${([^:}]+)[^}]*}`)
	yamlFile = []byte(re.ReplaceAllStringFunc(string(yamlFile), replaceByEnvVar))

	err = yaml.Unmarshal(yamlFile, &cfg)

	if err != nil {
		return err, Config{}
	}
	return nil, cfg
}

func replaceByEnvVar(match string) string {
	re := regexp.MustCompile(`\${([^:}]+)(?::([^}]*))?\s*}`)
	matchs := re.FindStringSubmatch(match)
	envValue := os.Getenv(matchs[1])

	if envValue != "" {
		return envValue
	}
	return matchs[2]
}
