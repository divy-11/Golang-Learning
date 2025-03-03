package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println(ptr)

	num := 17
	pointr := &num  // this pointr contains the refernce to num i.e. its address
	fmt.Println("The address of num :", pointr)
	fmt.Println("The value of num :", *pointr)

	*pointr-=11  //we did operation with pointer i.e. actual address of num Hence, changes occured.
	fmt.Println("Value of num :", num)
}
