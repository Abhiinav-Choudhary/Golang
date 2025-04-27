package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("welcome to learn about urls in golang")
	fmt.Println(myurl)
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}
	// fmt.Println(result.Host)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.User)

	qparams := result.Query()
	fmt.Printf("type of qparams is %T\n ", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	PartofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/login",
		RawPath: "user=Sanket",
	}

	AnotherUrl := PartofUrl.String()
	fmt.Println(AnotherUrl)
}
