package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"regexp"
)

func LoadConfig() Config {
	yamlFile, err := os.ReadFile("config/config.yaml")
	cfg := Config{}

	re := regexp.MustCompile(`\${([^:}]+)[^}]*}`)
	yamlFile = []byte(re.ReplaceAllStringFunc(string(yamlFile), replaceByEnvVar))

	err = yaml.Unmarshal(yamlFile, &cfg)

	if err != nil {
		panic(err)
	}

	return cfg
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
