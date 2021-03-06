package discord

import (
	"path/filepath"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type DiscordInfo struct {
	Email string
	Password string
	Token string
	BotId string
}

func ReadConfig(location string)(di DiscordInfo){
	filename, _ := filepath.Abs(location)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	di = DiscordInfo{}
	err = yaml.Unmarshal(yamlFile, &di)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return di
}