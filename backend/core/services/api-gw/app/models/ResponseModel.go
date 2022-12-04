package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type CompanyRequest struct {
	Name string `json:"name"`
}

type CabinetRequest struct {
	Id       string `json:"id"`
	Number   string `json:"number"`
	Size     string `json:"size"`
	Property string `json:"property"`
	Company  string `json:"company"`
}

type ScheduleRequest struct {
	CabinetId     string `json:"cabinet_id"`
	CabinetNumber string `json:"cabinet_number"`
	StartDate     string `json:"date_time_start"`
	EndDate       string `json:"date_time_end"`
	CompanyName   string `json:"company_name"`
}

func ResponseCompany(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST,OPTIONS")

	companyUrl := ConfigProcess().CompanyConfig.CompanyUrl

	var requestCompany CompanyRequest

	variables := mux.Vars(r)

	comapanyName, ok := variables["companyName"]

	if !ok {
		fmt.Println("Missing GET REST in parameters")

		err := json.NewDecoder(r.Body).Decode(&requestCompany)

		if err != nil {
			response, _ := json.Marshal(ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			Logger(1).Println(err)
			w.Write(response)
			return
		}
	} else {
		requestCompany.Name = comapanyName
	}

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	Logger(3).Println("Get request: ", r.Method, requestCompany)

	requestData := []byte(`{"name": "` + requestCompany.Name + `"}`)

	request, _ := http.NewRequest(r.Method, companyUrl, bytes.NewBuffer(requestData))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	Logger(4).Println("Send request: ", request)

	response, error := client.Do(request)

	if error != nil {
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad request!",
		})
		Logger(1).Println(error)
		w.Write(response)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Logger(1).Println(err)
	}

	Logger(4).Printf("Get response: %s\n", body)

	w.Write(body)
}

func ResponseCabinet(w http.ResponseWriter, r *http.Request) {

	cabinetUrl := ConfigProcess().CabinetConfig.CabinetUrl

	var requestCabinet CabinetRequest

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

	requestData := []byte(`{
		"id": "` + requestCabinet.Id + `",
		"number": "` + requestCabinet.Number + `",
		"size": "` + requestCabinet.Size + `",
		"property": "` + requestCabinet.Property + `",
		"company": "` + requestCabinet.Company + `"
	}`)

	request, _ := http.NewRequest(r.Method, cabinetUrl, bytes.NewBuffer(requestData))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	Logger(4).Println("Send request: ", request)

	response, error := client.Do(request)

	if error != nil {
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad request!",
		})
		Logger(1).Println(err)
		w.Write(response)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Logger(1).Println(err)
	}

	Logger(4).Printf("Get response: %s\n", body)

	w.Write(body)
}

func ResponseSchedule(w http.ResponseWriter, r *http.Request) {

	scheduleUrl := ConfigProcess().ScheduleConfig.ScheduleUrl

	var requestSchedule ScheduleRequest

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

	requestData := []byte(`{
		"cabinet_id": "` + requestSchedule.CabinetId + `",
		"cabinet_number": "` + requestSchedule.CabinetId + `",
		"date_time_start": "` + requestSchedule.StartDate + `",
		"date_time_end": "` + requestSchedule.EndDate + `",
		"company_name": "` + requestSchedule.CompanyName + `"
	}`)

	request, _ := http.NewRequest(r.Method, scheduleUrl, bytes.NewBuffer(requestData))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	Logger(4).Println("Send request: ", request)

	response, error := client.Do(request)

	if error != nil {
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad request!",
		})
		Logger(1).Println(err)
		w.Write(response)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Logger(1).Println(err)
	}

	Logger(4).Printf("Get response: %s\n", body)

	w.Write(body)
}
