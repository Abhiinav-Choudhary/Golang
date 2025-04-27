package main

import "fmt"

func main() {

	var fruitList [4]string
	fruitList[0] = "mango"
	fruitList[1] = "Apple"
	fruitList[3] = "Avocardo"

	fmt.Println("fruits in the given fruitlist are : ", fruitList)
	fmt.Println("length of fruitlist is : ", len(fruitList))

	var football = [4]string{"kipsta", "nivia", "cosco"}
	fmt.Println("football : ", football)
	fmt.Println("football : ", len(football))
}
