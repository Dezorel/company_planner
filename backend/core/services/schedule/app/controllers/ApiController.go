package controllers

import (
	"net/http"
	"schedule/app/models"
)

func ApiProcess() {

	http.HandleFunc("/api", models.Response)

	http.ListenAndServe(":"+models.ConfigProcess().ServiceConfig.Port, nil)
}
