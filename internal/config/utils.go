package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
)

func LoadConfig2(configType interface{}, configFileName string) interface{} {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	// FIXME call function to handle paths
	yamlFile, err := os.ReadFile(basepath + "/../../" + "config/" + configFileName)

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
