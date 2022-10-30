package controllers

import "company/app/models"

func ConfigProcess(configPath string) models.Config {

	var config models.Config

	models.ReadConfigFile(&config, configPath)

	return config
}
