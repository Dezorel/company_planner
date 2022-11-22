package controllers

import (
	"company/app/models"
	"net/http"
)

func ApiProcess() {

	http.HandleFunc("/test", models.Response)

	http.HandleFunc("/company", models.GetCompanyByName)

	http.ListenAndServe(":"+models.ConfigProcess().ServiceConfig.Port, nil)
}
