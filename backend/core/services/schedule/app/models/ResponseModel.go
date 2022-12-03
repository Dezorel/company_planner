package models

import (
	"encoding/json"
	"net/http"
)

type Company struct {
	Name string `json:"name"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, r *http.Request) {

	var requestCompany Company

	db := DBConnect()

	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&requestCompany)

	if err != nil {
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		w.Write(response)
		return
	}

	response, _ := json.Marshal(ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
	})

	switch r.Method {

	case "GET":
		//company, err := GetCompanyByName(requestCompany.Name)
		//
		//if err != nil {
		//	response, _ := json.Marshal(ErrorResponse{
		//		Status:  http.StatusBadRequest,
		//		Message: err.Error(),
		//	})
		//	w.Write(response)
		//	return
		//}
		//
		//response, _ = json.Marshal(company)

	case "POST":
		//company, err := CreateCompany(requestCompany.Name)

		//if err != nil {
		//	response, _ := json.Marshal(ErrorResponse{
		//		Status:  http.StatusBadRequest,
		//		Message: err.Error(),
		//	})
		//	w.Write(response)
		//	return
		//}
		//
		//response, _ = json.Marshal(company)
	default:

		w.WriteHeader(http.StatusMethodNotAllowed)

		response, _ = json.Marshal(ErrorResponse{
			Status:  http.StatusMethodNotAllowed,
			Message: "Message not allowed",
		})
	}

	w.Write(response)
}
