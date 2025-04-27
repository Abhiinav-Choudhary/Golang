package main

import (
	"fmt"
	"net/http"

	"github.com/Abhiinav-Choudhary/mongoApi/routes"
)

func main() {
	fmt.Println("MongoDB Api")
	r := routes.Routes()
	fmt.Println("server is stating ... ")
	http.ListenAndServe(":4000", r)
	fmt.Println("listening at port 4000")
}
