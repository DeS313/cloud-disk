package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	HOST string `json:"host"`
	PORT string `json:"port"`
	DB   struct {
		HOST     string `json:"host"`
		PORT     string `json:"port"`
		DATABASE string `json:"database"`
	}
	FilePath string `json:"filePath"`
}

func openConfigFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err, "ERROR READING FROM JSON FILE")
		return file, err
	}
	return file, err
}

func GetConfig(path string) (*Config, error) {
	var config Config
	file, err := openConfigFile(path)
	if err != nil {
		return &config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
		return &config, err
	}
	return &config, err
}
