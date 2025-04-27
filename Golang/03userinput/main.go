package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Name := "abhinav"
	fmt.Println("my name is ", Name)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("the lenght of my name is, ", input)

}
