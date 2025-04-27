package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status : ", response.StatusCode)
	fmt.Println("length : ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("byteCount is : ", byteCount)
	fmt.Println(responseString.String())

}

func PerformPostJsonRequest() {
	const url = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	"Name": "Abhinav,
	"phone":"redmi",
	"number":9179818502,
	"email":"abhinav123@gmail.com
	`)

	response, err := http. Post(myurl, "application/json", requestBody)	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
