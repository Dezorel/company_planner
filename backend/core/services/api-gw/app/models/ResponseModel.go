package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
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
	CabinetId string `json:"cabinet_id"`
	StartDate string `json:"date_time_start"`
	EndDate   string `json:"date_time_end"`
}

func ResponseCompany(w http.ResponseWriter, r *http.Request) {

	companyUrl := ConfigProcess().CompanyConfig.CompanyUrl

	var requestCompany CompanyRequest

	err := json.NewDecoder(r.Body).Decode(&requestCompany)

	if err != nil {
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		w.Write(response)
		return
	}

	requestData := []byte(`{"name": "` + requestCompany.Name + `"}`)

	request, _ := http.NewRequest(r.Method, companyUrl, bytes.NewBuffer(requestData))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad request!",
		})

		w.Write(response)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(body)
}

func ResponseCabinet(w http.ResponseWriter, r *http.Request) {

	cabinetUrl := ConfigProcess().CabinetConfig.CabinetUrl

	var requestCabinet CabinetRequest

	err := json.NewDecoder(r.Body).Decode(&requestCabinet)

	if err != nil {
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
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
	response, error := client.Do(request)
	if error != nil {
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad request!",
		})

		w.Write(response)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(body)
}

func ResponseSchedule(w http.ResponseWriter, r *http.Request) {

	scheduleUrl := ConfigProcess().ScheduleConfig.ScheduleUrl

	var requestSchedule ScheduleRequest

	err := json.NewDecoder(r.Body).Decode(&requestSchedule)

	if err != nil {
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		w.Write(response)
		return
	}

	requestData := []byte(`{
		"cabinet_id": "` + requestSchedule.CabinetId + `",
		"date_time_start": "` + requestSchedule.StartDate + `",
		"date_time_end": "` + requestSchedule.EndDate + `"
	}`)

	request, _ := http.NewRequest(r.Method, scheduleUrl, bytes.NewBuffer(requestData))

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Bad request!",
		})

		w.Write(response)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(body)
}
