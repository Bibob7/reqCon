package config

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)

func FromFile(filePath string) (*Config, error) {
	var config *Config
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if isYAML(filePath) {
		err = yaml.Unmarshal(fileContent, config)
		if err != nil {
			return nil, err
		}
	}
	if isJSON(filePath) {
		err = json.Unmarshal(fileContent, config)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

func isYAML(filePath string) bool {
	if strings.HasSuffix(filePath, ".yaml") {
		return true
	}
	if strings.HasSuffix(filePath, ".yml") {
		return true
	}
	return false
}

func isJSON(filePath string) bool {
	if strings.HasSuffix(filePath, ".json") {
		return true
	}
	return false
}
