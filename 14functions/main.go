package main

import "fmt"

func main() {
	fmt.Println("Hello from Main.")
	greet()
	fmt.Println("Simple Sum :", sumUp(2, 8))

	val1, val2 := proAdder(3, 2, 5, 7, 2)
	fmt.Println("Pro function :", val1, "Here is the string:", val2)
}

func greet() {
	fmt.Println("Namaste!!")
}

func sumUp(val1 int, val2 int) int { /*this int represent type we are returing ->Function Signatures*/
	return val1 + val2
}

//Pro Function -> can expect any number of inputs.
func proAdder(val ...int) (int, string) {
	tot := 0
	for _, v := range val {
		tot += v
	}
	return tot, "Hello trying something weird"
}
