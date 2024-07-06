package config

import (
	"encoding/json"
	"log"
	"os"
)

type JSONLoader struct{}

func (loader *JSONLoader) Load(fileName string, v interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close the file: %v", err)
		}
	}(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(v)
	return err
}

var _ ILoader = (*JSONLoader)(nil) // Ensure JSONConfigLoader implements ConfigLoader
