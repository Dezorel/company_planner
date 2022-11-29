package controllers

import (
	"cabinet/app/models"
	"net/http"
)

func ApiProcess() {

	http.HandleFunc("/api", models.Response)

	http.ListenAndServe(":"+models.ConfigProcess().ServiceConfig.Port, nil)
}
