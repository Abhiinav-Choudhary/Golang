package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome in the golang code to learn about files")
	content := "hello welcome everyone in my amazing golang codeEnv"
	file, err := os.Create("./coding.txt")
	// if err != nil {
	// 	fmt.Println("404 not found", err)
	// }
	checkErr(err)
	length, err := io.WriteString(file, content)
	checkErr(err)
	// if err != nil {
	// 	fmt.Println("404 not found", err)
	// }
	fmt.Println("length : ", length)
	defer file.Close()
	Readfile("./coding.txt")
}

func Readfile(file string) {
	databytes, err := ioutil.ReadFile(file)
	checkErr(err)
	// if err != nil {
	// 	fmt.Println("404 not found", err)
	// }
	fmt.Println(string(databytes))
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("404 not found", err)
	}
}
