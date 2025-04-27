package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to study time of golang")

	Showtime := time.Now()
	fmt.Println(Showtime)
	fmt.Println(Showtime.Format("01-02-2006 15-04-05 monday"))

	Createtime := time.Date(2025, time.August, 25, 20, 10, 0, 0, time.UTC)
	fmt.Println(Createtime.Format("01-02-2006 monday"))

}
