package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate us")

	input, _ := reader.ReadString('\n') // comma ok syntax
	/*
		reader.ReadString gives two values result and error (input, err)
		in most of the cases we dont care about errors in Go so, use _ at place of err
		In few cases, we care only about error and not the input -> (_ , err)
	*/

	fmt.Println("Thanks for rating us, ", input)
	fmt.Printf("Type of input is: %T", input)
}
