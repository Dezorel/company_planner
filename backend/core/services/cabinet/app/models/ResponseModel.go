package models

import (
	"encoding/json"
	"net/http"
)

type Cabinet struct {
	Id       string `json:"id"`
	Number   string `json:"number"`
	Size     string `json:"size"`
	Property string `json:"property"`
	Company  string `json:"company"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, r *http.Request) {

	var requestCabinet Cabinet

	db := DBConnect()

	defer db.Close()

	err := json.NewDecoder(r.Body).Decode(&requestCabinet)

	Logger(3).Println("Get request: ", r.Method, requestCabinet)

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
		if requestCabinet.Id == "" {
			cabinet, err := GetCabinetsByCompany(requestCabinet.Company)

			if err != nil {
				response, _ := json.Marshal(ErrorResponse{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				})
				Logger(1).Println(err)
				w.Write(response)
				return
			}

			response, _ = json.Marshal(cabinet)
		} else {
			cabinet, err := GetCabinetInfoById(requestCabinet.Id)

			if err != nil {
				response, _ := json.Marshal(ErrorResponse{
					Status:  http.StatusBadRequest,
					Message: err.Error(),
				})
				Logger(1).Println(err)
				w.Write(response)
				return
			}

			response, _ = json.Marshal(cabinet)
		}
	case "POST":
		cabinet, err := CreateCabinet(requestCabinet.Company, requestCabinet.Number, requestCabinet.Size, requestCabinet.Property)

		if err != nil {
			response, _ := json.Marshal(ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			Logger(1).Println(err)
			w.Write(response)
			return
		}

		response, _ = json.Marshal(cabinet)
	default:

		w.WriteHeader(http.StatusMethodNotAllowed)

		response, _ = json.Marshal(ErrorResponse{
			Status:  http.StatusMethodNotAllowed,
			Message: "Message not allowed",
		})
	}

	w.Write(response)
}
