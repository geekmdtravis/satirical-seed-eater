package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Version string `json:"version"`
}

func GetConfig() Config {
	var config Config
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(file, &config)
	return config
}