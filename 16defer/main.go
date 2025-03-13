package main

import "fmt"

func main() {
	//defer holds the line in a LIFO stack and executes the stack after whole normal code.
	defer fmt.Println("hey1")
	fmt.Println("Hey2")
	defer fmt.Println("hey3")
	defer fmt.Println("hey4")
	fmt.Println("hey5")
}
