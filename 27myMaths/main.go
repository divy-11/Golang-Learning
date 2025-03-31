package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	var val1 int = 3
	var val2 float64 = 4.5
	fmt.Println("The sum is :", val2+float64(val1))

	//Random Number
	//OLD THING - DEPRICATED WITH UPDATE :: rand.Seed(321) -> seed determines the random number generated until seed is same Random number will remain same

	/*
	   1. MATH RANDOM
	   fmt.Println(rand.Intn(6)) //In modern Languages the range is always EXCLUSIVE (0-5)
	*/

	// 2. Cryptographic Way
	randNum, err := rand.Int(rand.Reader, big.NewInt(6))
	if err != nil {
		panic(err)
	}
	fmt.Println("Random Number :", randNum)
}
