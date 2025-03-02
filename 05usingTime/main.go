package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	// to format we need to use these specified values only.
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("Monday"))
	fmt.Println(presentTime.Format("15:04:05"))

	createdDate := time.Date(2023, time.August, 05, 23, 59, 59, 900000000, time.UTC)
	fmt.Println("Created Time : ", createdDate.Format("01-02-2006 Monday 15:04:05"))
}
