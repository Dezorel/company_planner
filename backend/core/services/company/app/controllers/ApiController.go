package controllers

import (
	"company/app/models"
	"net/http"
)

func Process(bindAddr string) {

	http.HandleFunc("/api", models.Response)

	http.ListenAndServe(":"+bindAddr, nil)
}
