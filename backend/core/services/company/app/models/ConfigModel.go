package models

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Service struct {
	Port string `yaml:"port"`
}

type Config struct {
	Service `yaml:"service"`
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
