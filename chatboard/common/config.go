package common

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Adress    string `json:"adress"`
	DBName    string `json:"dbname"`
	LogToFile bool   `json:"log_to_file"`
}

var config Configuration

func LoadConfig() (Configuration, error) {
	confName := "config.json"
	file, err := os.Open(confName)
	if err != nil {
		log.Fatalln("could not open config file.")
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("could not decode config file.")
	}
	return config, err
}
