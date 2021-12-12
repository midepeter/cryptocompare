package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

var configPath string = "config.json"

type Config struct {
	Db  DB  `json:"db"`
	Key Key `json:"key"`
}
type DB struct {
	Password string `json:"password"`
	User     string `json:"user"`
	Name     string `json:"dbname"`
	Port     string `json:"port"`
	Host     string `json:"host"`
}

type Key struct {
	Apikey string `json:"apikey"`
}

func GetConfig() Config {
	path, _ := filepath.Abs(configPath)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("Invalid config path : "+configPath, err)
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalln("Failed unmarshal config ", err)
	}
	return config
}
