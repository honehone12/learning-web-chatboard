package common

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Adress string
	DBName string
	DBSync bool
}

var config Configuration

func LoadConfig() (Configuration, error) {
	confName := "config.json"
	file, err := os.Open(confName)
	if err != nil {
		log.Fatalln("could not open config file.")
		return config, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("could not decode config file.")
	}
	return config, err
}
