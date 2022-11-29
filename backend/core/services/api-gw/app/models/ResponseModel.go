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
