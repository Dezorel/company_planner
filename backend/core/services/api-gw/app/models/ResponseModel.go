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

func ResponseCompany(w http.ResponseWriter, r *http.Request) {

	companyUrl := "http://company:13070/api"

	data := []byte(`{"name": "test"}`)

	request, _ := http.NewRequest(r.Method, companyUrl, bytes.NewBuffer(data))

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
