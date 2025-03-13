package main

import "fmt"

func main() {
	fmt.Println("No Inheritence in Golang | No Super or parent")
	deva := User{"Dev", "Dev@gmail.com", 20, true}
	fmt.Println(deva)
	fmt.Printf("Value: %+v\n", deva)
	fmt.Printf("User is %v and is %v years old.", deva.Name, deva.Age)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Active bool
	//one int  this property wont be exportable as first alpha is small
}
