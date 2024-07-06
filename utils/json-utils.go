package utils

import (
	"encoding/json"
	"os"
)

type JSONLoader struct{}

func (loader *JSONLoader) LoadJson(fileName string, v interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(v)
	return err
}
