package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var url = "https://jsonplaceholder.typicode.com/posts"

func main() {

	// dataGET, err := GET(url)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("+%v\n", string(dataGET))

	dataPOST, err := POST(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("+%v\n", string(dataPOST))
}

func GET(url string) ([]byte, error) {
	data, err := req(url, http.MethodGet, nil)
	if err != nil {
		panic(err)
	}

	return data, nil
}

func POST(url string) ([]byte, error) {
	var payload = map[string]interface{}{
		"userId": 1,
		"title":  "MNC B",
		"body":   "body body",
	}

	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	data, err := req(url, http.MethodPost, bytePayload)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func req(url string, method string, data []byte) ([]byte, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var dataInterface interface{}
	err = json.NewDecoder(resp.Body).Decode(&dataInterface)
	if err != nil {
		return nil, err
	}

	dataByte, err := json.Marshal(dataInterface)
	if err != nil {
		return nil, err
	}

	return dataByte, nil
}
