package models

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const CONFIG_PATH = "./config/config.yaml"

type ServiceConfig struct {
	Port string `yaml:"port"`
}

type CompanyConfig struct {
	CompanyUrl string `yaml:"url"`
}

type CabinetConfig struct {
	CabinetUrl string `yaml:"url"`
}

type Config struct {
	ServiceConfig `yaml:"service"`
	CompanyConfig `yaml:"company"`
	CabinetConfig `yaml:"cabinet"`
}

func ReadConfigFile(cfg *Config, configPath string) {
	f, err := os.Open(configPath)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func ConfigProcess() Config {

	var config Config

	ReadConfigFile(&config, CONFIG_PATH)

	return config
}
