package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rate us for the service : ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating !", input)
	//input is of type STRING. We want to be var. We want to do some opts

	numInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	/*
		this also throughs two values hence use (val,err)

		WrongWayInp , err1 := strconv.ParseFloat(input, 64)
		if we use directly input, input="4\n" i.e contains space. Need to be trimmed.
	*/

	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("Converted rating: (", numInput/2,"/5 )")
	}

}
