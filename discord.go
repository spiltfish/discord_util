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

func Read_config()(di DiscordInfo){
	filename, _ := filepath.Abs("./discord_config.yml")
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