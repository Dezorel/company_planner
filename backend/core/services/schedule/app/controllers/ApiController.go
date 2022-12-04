package controllers

import (
	"net/http"
	"schedule/app/models"
)

func ApiProcess() {

	http.HandleFunc("/api", models.Response)

	models.Logger(3).Println("Schedule API started")

	http.ListenAndServe(":"+models.ConfigProcess().ServiceConfig.Port, nil)
}
