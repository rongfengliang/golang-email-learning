package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config cmp sync config
type Config struct {
	Email struct {
		ServerHost string `yaml:"serverhost"`
		ServerPort int    `yaml:"serverport"`
		FromEmail  string `yaml:"fromemail"`
		FromPasswd string `yaml:"from_passwd"`
	} `yaml:"email"`
	Template struct {
		EmailTemplate string `yaml:"email"`
	} `yaml:"template"`
}

// Default config path is local with name config.yaml

// New get sync config
func New() Config {
	config := Config{}
	bytes, err := ioutil.ReadFile("config.yaml")
	log.Printf("%s", bytes)
	if err != nil {
		log.Fatalln("read config error: ", err.Error())
	}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalln("Unmarshal config error: ", err.Error())
	}
	return config
}
