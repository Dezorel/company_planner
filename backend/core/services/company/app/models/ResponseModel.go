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

func GetCompanyByName(w http.ResponseWriter, r *http.Request) {

	var company Company

	err := json.NewDecoder(r.Body).Decode(&company)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db := DBConnect()

	defer db.Close()

	query := "SELECT 1 as result FROM `Companies` WHERE company_name = '" + company.Name + "'"

	resultQuery := DBQueryRow(db, query)

	if resultQuery.Result != "1" {
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad Request",
		})

		w.Write(response)

	} else {
		response, _ := json.Marshal(Company{
			Name: company.Name,
		})

		w.Write(response)
	}

}
