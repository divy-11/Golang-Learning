package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate us")
	input, _ := reader.ReadString('\n')   // comma ok syntax
	fmt.Println("Thanks for rating us, ", input)
	fmt.Printf("Type of input is: %T", input)
}
