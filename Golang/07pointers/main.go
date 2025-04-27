package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println(ptr)

	Number := 34
	p := &Number
	fmt.Println("val of p : ", p)
	fmt.Println("val of p : ", *p)
	fmt.Println("val of p : ", *p+2)
}
