package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type mongoConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
}

type globalConfig struct {
	Mongo mongoConfig `yaml:"mongo"`
}

var GConfig = &globalConfig{}

func init() {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, GConfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
