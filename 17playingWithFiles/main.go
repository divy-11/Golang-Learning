package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "This is testing on creating a .txt file"

	file, err := os.Create("./test.txt") //to create file
	errHandle(err)

	len, error := io.WriteString(file, content) //to add content
	errHandle(error)
	fmt.Println("Length of file:", len)
	defer file.Close()

	ReadF("./test.txt")
}

func ReadF(file string) {
	dataBytes, err := os.ReadFile(file) //data read here is in BYTES FORMAT and not string
	errHandle(err)
	fmt.Println("Data inside file is :", string(dataBytes))
}

func errHandle(e error) {
	if e != nil {
		panic(e)
	}
}