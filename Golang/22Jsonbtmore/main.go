package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"studentName"`
	Rollno   string `json:"rollno"`
	Id       string `json:"id"`
	Password string `json:"-"`
	Age      int
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	DecoderJson()
}

func EncodeJson() {
	fmt.Println("learn about json and structs")
	StudentInfo := []course{
		{"Abhinav", "23ESKCX004", "B230439", "Abhinav@133", 20, []string{"web-dev", "js"}},
		{"Akshat", "23ESKCX008", "B230432", "Akshat@123", 19, []string{"web-dev", "js"}},
		{"Aishnai", "23ESKCX006", "B231342", "Aishani@143", 18, nil},
	}

	Finaljson, err := json.MarshalIndent(StudentInfo, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", Finaljson)
}

func DecoderJson() {
	ProductInfo := []byte(`{
	 "ProductName" : "laptop",
	 "productId " :"34#D",
	 "prouctPrice": 456789,
	 "warrenty" : "1 year"
	 }`)

	var Product course
	var validJson = json.Valid(ProductInfo)
	if validJson {
		fmt.Println("The json is valid")
		json.Unmarshal(ProductInfo, &Product)
		fmt.Printf("%#v\n", Product)
	} else {
		fmt.Println("json is not valid")
	}

	var OtherPro map[string]interface{}
	json.Unmarshal(ProductInfo, &OtherPro)
	fmt.Printf("%#v\n", OtherPro)

	for k, v := range OtherPro {
		fmt.Printf("key is %v\n and val is %v\n and type is : %T\n", k, v, v)
	}
}
