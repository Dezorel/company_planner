package controllers

import (
	"api-gw/app/models"
	"net/http"
)

func GateWayProcess() {
	http.HandleFunc("/api/company", models.ResponseCompany)

	http.HandleFunc("/api/cabinet", models.ResponseCabinet)

	http.HandleFunc("/api/schedule", models.ResponseSchedule)

	models.Logger(3).Println("Gateway started")

	http.ListenAndServe(":"+models.ConfigProcess().ServiceConfig.Port, nil)
}
