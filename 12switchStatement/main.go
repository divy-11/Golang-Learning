package main

import (
	"fmt"
	"math/rand"
)

func main() {
	diceRoll := rand.Intn(6) + 1 //as funct gives number from 0-5
	switch diceRoll {
	case 1:
		fmt.Println("You can open.")
	case 2:
		fmt.Println("Move two steps.")
	case 3:
		fmt.Println("Move three steps.")
	case 4:
		fmt.Println("Move four steps.")
	case 5:
		fmt.Println("Move five steps.")
	case 6:
		fmt.Println("Move six steps and roll again.")
		fallthrough //after running this case will move to next case even if default and run that too

	default: //any value other than cases
		fmt.Println("What was that!!")
	}
}
