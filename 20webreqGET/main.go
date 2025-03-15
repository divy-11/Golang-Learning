package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	const myURL string = "http://localhost:8000/"
	resp, err := http.Get(myURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status Code :", resp.StatusCode)
	fmt.Println("Length of Response: ", resp.ContentLength)
	defer resp.Body.Close()

	contentByte, _ := io.ReadAll(resp.Body)
	fmt.Println(string(contentByte))

	//different approach of stringifying
	var respData strings.Builder
	
	byteLen, _ := respData.Write(contentByte) //saves the info in respData variable permanently.
	fmt.Println(byteLen) //returns byte length

	fmt.Println(respData.String())
}
