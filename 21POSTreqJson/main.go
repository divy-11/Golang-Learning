package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	const myURL string = "http://localhost:8000/post"

	contentJSON := strings.NewReader(`
		{
		"coursename":"Cohort 2.0",
		"price":"2999",
		"country":"India"	
		}
	`)
	resp, err := http.Post(myURL, "application/json", contentJSON)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)
	fmt.Println(string(content))
}
