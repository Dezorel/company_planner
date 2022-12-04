package models

import (
	"gopkg.in/yaml.v3"
	"os"
)

const CONFIG_PATH = "./config/config.yaml"

type ServiceConfig struct {
	Port     string `yaml:"port"`
	LogLevel uint8  `yaml:"logLevel"`
}

type CompanyConfig struct {
	CompanyUrl string `yaml:"url"`
}

type CabinetConfig struct {
	CabinetUrl string `yaml:"url"`
}

type ScheduleConfig struct {
	ScheduleUrl string `yaml:"url"`
}

type Config struct {
	ServiceConfig  `yaml:"service"`
	CompanyConfig  `yaml:"company"`
	CabinetConfig  `yaml:"cabinet"`
	ScheduleConfig `yaml:"schedule"`
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
	Logger(1).Println(err)
	os.Exit(2)
}

func ConfigProcess() Config {

	var config Config

	ReadConfigFile(&config, CONFIG_PATH)

	return config
}
