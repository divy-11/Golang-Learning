package main

import "fmt"

func main() {
	cities := []string{"Dhanbad", "Banaras", "Lucknow", "Jalandhar"}
	for d := 0; d < len(cities); d++ {
		fmt.Println(cities[d])
	}
	for i := range cities {
		fmt.Println(cities[i])
	}

	for ind, val := range cities { //for each loop
		fmt.Printf("The index is %v and the value is %v \n", ind, val)
	}

	valInt := 0
	for valInt < 10 {
		if valInt == 2 {
			//we cant simply continue. it wont move!
			valInt += 2 //we need to provide the value after continue. how it would change.
			continue
		}
		if valInt == 5 {
			goto devaStatement
		}
		if valInt == 8 {
			break
		}
		fmt.Println("Value :", valInt)
		valInt++
	}

devaStatement: //goto statement should always be after the place of call
	fmt.Println("Hello from GOTO Example.")

}
