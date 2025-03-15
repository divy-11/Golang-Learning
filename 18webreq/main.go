package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://lco.dev"
func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", resp)
	defer resp.Body.Close() //its our resp to close the connection in the end.
	contentBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(contentBytes))
}
