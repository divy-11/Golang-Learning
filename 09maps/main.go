package main

import "fmt"

func main() {
	//maps -> hash Table -> key value pair

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["TS"] = "Typescript"
	languages["RB"] = "Ruby"
	fmt.Println("Showing :", languages)

	delete(languages, "RB")
	fmt.Println("After deleting RB :", languages)

	for key, val := range languages {  // we can use ( _ , val ) if we dont wanna print key
		fmt.Printf("The key is %v and the value is %v \n", key, val)
	}
}
