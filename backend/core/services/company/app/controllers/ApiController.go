package controllers

import (
	"company/app/models"
	"net/http"
)

func ApiProcess(bindAddr string) {

	http.HandleFunc("/api", models.Response)

	http.ListenAndServe(":"+bindAddr, nil)
}
