package main

import "fmt"

func main() {
	fmt.Println("No Inheritence in Golang | No Super or parent")
	deva := User{"Dev", "Dev@gmail.com", 20, true}
	fmt.Println(deva)
	fmt.Printf("Value: %+v\n", deva)

	deva.GetStatus()

	// fmt.Printf("User is %v and is %v years old.\n", deva.Name, deva.Age)
	deva.GetAge() //using method.

	deva.ChangeStatus()
	deva.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Age    int
	Active bool
}

//Method
func (u User) GetStatus() { // (input inputType)
	fmt.Printf("Status of %v : %v \n", u.Name, u.Active)
}

func (u User) GetAge() {
	fmt.Printf("User is %v and is %v years old.\n", u.Name, u.Age)
}

func (u User) ChangeStatus() {
	//this u passed in methods is a copy and not the original. Manipulation cant be done like this
	u.Active = !u.Active
	u.GetStatus()
}
