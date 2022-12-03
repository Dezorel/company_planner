package controllers

import (
	"company/app/models"
	"net/http"
)

func ApiProcess() {

	http.HandleFunc("/api", models.Response)

	models.Logger(3).Println("Company API started")

	http.ListenAndServe(":"+models.ConfigProcess().ServiceConfig.Port, nil)
}
