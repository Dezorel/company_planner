package controllers

import (
	"api-gw/app/models"
	"net/http"
)

func GateWayProcess() {
	http.HandleFunc("/api/company", models.ResponseCompany)

	http.HandleFunc("/api/cabinet", models.ResponseCabinet)

	http.ListenAndServe(":"+models.ConfigProcess().ServiceConfig.Port, nil)
}
