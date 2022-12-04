package controllers

import (
	"api-gw/app/models"
	"github.com/gorilla/mux"
	"net/http"
)

func GateWayProcess() {
	router := mux.NewRouter()

	router.HandleFunc("/api/company/{companyName}", models.ResponseCompany)

	router.HandleFunc("/api/cabinet/{companyName}", models.ResponseCabinet)

	router.HandleFunc("/api/schedule/{companyName}", models.ResponseSchedule)

	router.HandleFunc("/api/company", models.ResponseCompany)

	router.HandleFunc("/api/cabinet", models.ResponseCabinet)

	router.HandleFunc("/api/schedule", models.ResponseSchedule)

	models.Logger(3).Println("Gateway started")

	http.ListenAndServe(":"+models.ConfigProcess().ServiceConfig.Port, router)
}
