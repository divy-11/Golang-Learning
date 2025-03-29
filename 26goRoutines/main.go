package main

import (
	"fmt"
	"net/http"
	"sync"
)

var waitGrp sync.WaitGroup // will wait has three fns: Add, Wait and Done

func main() {
	websiteList := []string{
		"https://google.com",
		"https://x.com",
		"https://github.com",
		"https://realmadrid.com",
	}

	//Go Routines
	for _, w := range websiteList {
		go getStatus(w)
		waitGrp.Add(1) //add the number of go routines, here we are sending one at a time.
	}

	waitGrp.Wait() //this wont let main fn end until all go routines are done.

}

func getStatus(s string) {
	resp, err := http.Get(s)
	if err != nil {
		fmt.Println("Error occur.")
	} else {
		fmt.Printf("%d status code for %s\n", resp.StatusCode, s)
	}
	defer waitGrp.Done() //will tell that Go routine is finished here. Always keep at the end(defer).
}
