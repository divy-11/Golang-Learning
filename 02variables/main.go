package main

import "fmt"

const GlobalVar string ="hello first alpha caps means global"

func main() {
	var username string = "Devv"
	fmt.Println(username)
	fmt.Printf("Variable Type : %T \n", username)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable Type : %T \n", smallVal)

	var floatVal float32 = 256.632742538
	fmt.Println(floatVal)
	fmt.Printf("Variable Type : %T \n", floatVal)

	var num=2937
	fmt.Println(num)

	numofuser := 3883.03
	fmt.Printf("Variable Type : %T \n", numofuser)
	fmt.Println(numofuser)

	fmt.Println(GlobalVar)
}
