package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	getRequest()
	postRequest()
}

func getRequest() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func postRequest() {
	data := url.Values{"name": {"John Doe"}}

	resp, err := http.PostForm("https://jsonplaceholder.typicode.com/posts", data)

	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
