package models

import (
	"encoding/json"
	"net/http"
)

type Company struct {
	Name string `json:"name"`
	Date string `json:"datetime"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, r *http.Request) {

	db := DBConnect()

	defer db.Close()

	query := "SELECT NOW()"

	resultQuery := DBQueryRow(db, query)

	company := Company{
		Name: "UTM",
		Date: resultQuery.Result,
	}

	response, _ := json.Marshal(ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
	})

	switch r.Method {

	case "GET":

		response, _ = json.Marshal(company)

	default:

		w.WriteHeader(http.StatusMethodNotAllowed)

		response, _ = json.Marshal(ErrorResponse{
			Status:  http.StatusMethodNotAllowed,
			Message: "Message not allowed",
		})
	}

	w.Write(response)
}
