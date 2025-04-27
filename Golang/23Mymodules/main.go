package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello mod learners")
	greet()
	r := mux.NewRouter()
	r.HandleFunc("/", server).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}
func greet() {
	fmt.Println("hello coders")
}

func server(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Hello welcome to golang </h1>`))
}
