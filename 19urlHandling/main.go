package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://localhost:3000/courses?name=cohort2&userid=lAish17"

func main() {
	fmt.Println(myUrl)

	//parsing the url
	resp, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Host)
	fmt.Println(resp.Path)
	fmt.Println(resp.Scheme)
	fmt.Println(resp.RawQuery)
	fmt.Println(resp.Port())

	qparams := resp.Query() //better way of storing (key,value)pair
	fmt.Println(qparams)
	fmt.Printf("The type of query is : %T\n", qparams)

	for _, v := range qparams {
		fmt.Println(v)
	}

	//constructing URL
	newURL := &url.URL{
		Scheme:  "https",
		Host:    "localhost:8000",
		Path:    "test",
		RawPath: "app=krishiSakha",
	}
	fmt.Println("Final URL : ", newURL.String())
}
