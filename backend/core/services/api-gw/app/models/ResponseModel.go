package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://127.0.0.1")

	switch r.Method {
	case "GET":
		resp, err = http.Get("https://jsonplaceholder.typicode.com/posts")

	case "POST":
		data := url.Values{"name": {"John Doe"}}
		resp, err = http.PostForm("https://jsonplaceholder.typicode.com/posts", data)

	default:
		response, _ := json.Marshal(ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Method not allowed!",
		})

		w.Write(response)

		return
	}

	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	w.Write(body)
}
