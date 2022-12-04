package controllers

import (
	"cabinet/app/models"
	"net/http"
)

func ApiProcess() {

	http.HandleFunc("/api", models.Response)

	models.Logger(3).Println("Cabinet API started")

	http.ListenAndServe(":"+models.ConfigProcess().ServiceConfig.Port, nil)
}
