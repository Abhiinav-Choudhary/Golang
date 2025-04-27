package main

import (
	"fmt"
	"sort"
)

func main() {
	var slices = []int{1, 2, 3, 4, 4}
	slices = append(slices[:3])
	fmt.Println(slices)
	fmt.Printf("type of slices is %T\n ", slices)

	lang := []string{"js", "rb"}
	// lang[1] = "java"
	// lang[2] = "ruby"
	// lang[3] = "python"
	fmt.Println(append(lang[1:], "py"))

	highschools := make([]string, 4)
	highschools[0] = "DPS"
	highschools[1] = "GVM"
	highschools[2] = "GYANGANGA"
	highschools[3] = "NHGOEL"
	// highschools[4] = "BRIGHTOON"
	fmt.Println(highschools)
	fmt.Println(append(highschools, "morningstar", "Angelmontesari"))

	highscores := make([]int, 5)
	highscores[0] = 99
	highscores[1] = 67
	highscores[2] = 56
	highscores[3] = 72
	highscores[4] = 93
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

	var vegan = []string{"cabbage", "potato", "tomato", "latus", "ladyfinger"}
	fmt.Println("list of vegan items are : ", vegan)
	fmt.Println("after removing tomato list will be : ", append(vegan[:2], vegan[3:]...))
}
