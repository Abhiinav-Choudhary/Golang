package main

import "fmt"

var bcrypt string = "password"

func main() {
	var name string = "Abhinav"
	fmt.Println(name)

	var id uint32 = 2565632123
	fmt.Println(id)
	fmt.Printf("type of id: %T\n", id)

	var isLogged bool = false
	fmt.Println(isLogged)
	fmt.Printf("type of if: %T\n", isLogged)

	JWT := "jsonwebtoken"
	fmt.Println(JWT)
	fmt.Printf("type of if: %T\n", JWT)

	fmt.Println(bcrypt)

}
