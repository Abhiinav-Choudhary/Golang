package main

import "fmt"

func main() {
	lang := make(map[string]string)
	lang["js"] = "javascript"
	lang["rb"] = "ruby"
	lang["py"] = "python"
	fmt.Println("list of languages", lang)

	fmt.Println("rb sort of : ", lang["rb"])
	delete(lang, "js")
	fmt.Println("lang after abstraction of js : ", lang)

	// loops in golang are interesting

	for _, value := range lang {
		fmt.Printf("For key v  , val is %v\n", value)
	}

}
