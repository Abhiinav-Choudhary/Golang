package main

import "fmt"

func main() {
	fmt.Println("welcome in the golang to learn about ifelse")

	lowerCount := 10

	if lowerCount < 10 {
		fmt.Println("invalid lowercount")
	} else if lowerCount > 10 {
		fmt.Println("it is valid ")
	} else {
		lowerCount++
		fmt.Println("now it is valid")
		fmt.Println(lowerCount)
	}

	if nums := 13; nums < 10 {
		fmt.Println("nums is less than 10")
	} else {
		fmt.Println("nums is greater than 10")
	}

}
