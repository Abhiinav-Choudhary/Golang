package main

import "fmt"

func main() {
	fmt.Println("welcome to learn structres in golang")
	Abhinav := User{"Abhinav", "kumar", "Abhi@gmail.com", "Sanket@123", 1982739409}
	fmt.Printf("Abhinav details are %+v :", Abhinav)
	fmt.Printf("Name is %v and Email is %v ", Abhinav.FirstName, Abhinav.Email)
}

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Number    int
}
