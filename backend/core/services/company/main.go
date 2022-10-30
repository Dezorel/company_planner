package main

import "company/app/controllers"

const CONFIG_PATH = "./config/config.yaml"

func main() {

	config := controllers.ConfigProcess(CONFIG_PATH)

	controllers.ApiProcess(config.Service.Port)
}
