package controllers

import (
	"api-gw/app/models"
	"net/http"
)

func GateWayProcess() {
	http.HandleFunc("/api", models.Response)

	http.ListenAndServe(":"+models.ConfigProcess().ServiceConfig.Port, nil)
}
