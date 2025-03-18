package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	reqUsingJSON()
	reqUsingFormData()
}

func reqUsingJSON() {
	const myURL string = "http://localhost:8000/post"

	contentJSON := strings.NewReader(` //simple way to make JSON data
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

func reqUsingFormData() {
	const myURL string = "http://localhost:8000/postform"

	data := url.Values{} //key,value pair
	data.Add("firstName", "Dev")
	data.Add("Country", "India")
	data.Add("Email", "devsangwan@gmail.com")
	fmt.Println(data) 
	fmt.Printf("Type of data is %T\n",data) //url.Values
	resp, _ := http.PostForm(myURL, data)
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)
	fmt.Println(string(content))
}
