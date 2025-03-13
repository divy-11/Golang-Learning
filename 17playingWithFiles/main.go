package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is testing on creating a .txt file"
	file, err := os.Create("./test.txt")
	errHandle(err)
	len, error := io.WriteString(file, content)
	errHandle(error)
	fmt.Println("Length of file:", len)
	defer file.Close()
	ReadFile("./test.txt")
}

func errHandle(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(file string) {
	dataBytes, err := ioutil.ReadFile(file) //data read here is in BYTES FORMAT and not string
	errHandle(err)
	fmt.Println("Data inside file is :", string(dataBytes))
}
