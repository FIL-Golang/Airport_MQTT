package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"regexp"
)

func LoadConfig(configType interface{}, configFileName string) interface{} {
	yamlFile, err := os.ReadFile("config/" + configFileName)

	re := regexp.MustCompile(`\${([^:}]+)[^}]*}`)
	yamlFile = []byte(re.ReplaceAllStringFunc(string(yamlFile), replaceByEnvVar))

	err = yaml.Unmarshal(yamlFile, configType)

	if err != nil {
		panic(err)
	}

	return configType
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
