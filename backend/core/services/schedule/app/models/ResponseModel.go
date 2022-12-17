package models

import (
	"encoding/json"
	"net/http"
)

type Schedule struct {
	CabinetId     string `json:"cabinet_id"`
	CabinetNumber string `json:"cabinet_number"`
	StartDate     string `json:"date_time_start"`
	EndDate       string `json:"date_time_end"`
	CompanyName   string `json:"company_name"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, r *http.Request) {

	var requestSchedule Schedule

	db := DBConnect()

	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&requestSchedule)

	Logger(3).Println("Get request: ", r.Method, requestSchedule)

	if err != nil {
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		Logger(1).Println(err)
		w.Write(response)
		return
	}

	response, _ := json.Marshal(ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
	})

	switch r.Method {

	case "GET":
		schedules, err := GetScheduleByCompanyId(requestSchedule.CompanyName)

		if err != nil {
			response, _ := json.Marshal(ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			Logger(1).Println(err)
			w.Write(response)
			return
		}

		response, _ = json.Marshal(schedules)

	case "POST":
		schedules, err := CreateSchedule(requestSchedule.CompanyName, requestSchedule.StartDate, requestSchedule.EndDate, requestSchedule.CabinetNumber)

		if err != nil {
			response, _ := json.Marshal(ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			Logger(1).Println(err)
			w.Write(response)
			return
		}

		response, _ = json.Marshal(schedules)
	default:

		w.WriteHeader(http.StatusMethodNotAllowed)

		response, _ = json.Marshal(ErrorResponse{
			Status:  http.StatusMethodNotAllowed,
			Message: "Message not allowed",
		})
	}

	w.Write(response)
}
