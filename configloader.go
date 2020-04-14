package configloader

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfigYaml(v interface{}, fileName string) (interface{}, error) {
	configFile, error := os.Open(fileName)
	if error != nil {
		return nil, errors.New("config file not found")
	}
	defer configFile.Close()

	t := yaml.NewDecoder(configFile)
	if tError := t.Decode(v); tError != nil {
		return nil, tError
	}

	return v, nil
}
