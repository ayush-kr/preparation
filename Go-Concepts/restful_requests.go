package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 5 * time.Second,
	}
}

func main() {

	//GET request
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		log.Println("Error in Get!", err.Error())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error in Get!", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Result: ", string(body))

	//POST request
	jsonData := map[string]string{"id": "124", "value": "hells"}
	jsonByte, _ := json.Marshal(jsonData)

	postRequest, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonByte))
	if err != nil {
		log.Println(err.Error())
		return
	}
	postRequest.Header.Add("Content-Type", "application/json")
	postResp, err := client.Do(postRequest)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer postResp.Body.Close()

	postRespBody, err := ioutil.ReadAll(postResp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Post Result:", string(postRespBody))
}
