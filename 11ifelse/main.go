package main

import "fmt"

func main() {

	var isCnt int = 23
	if isCnt%2 == 0 {
		fmt.Printf("Is Even")
	}else if isCnt%2==1{
		fmt.Printf("Is Odd")
	}else{
		fmt.Printf("NaN")
	}
}
