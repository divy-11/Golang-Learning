package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Peach", "Mango"} //in slice no need to pre-define len
	fmt.Printf("Type of :%T \n", fruitList)
	fmt.Println("List :", fruitList)

	fruitList = append(fruitList, "Banana", "Orange")
	fmt.Println("List after appending :", fruitList)

	fruitList = append(fruitList[1:4]) //this will slice it from index 1 to 3 (doesnt 4th one i.e. last range is not INCLUSIVE)
	fmt.Println("List after slicing :", fruitList)

	// fruitList = append(fruitList[:3])  -> [Apple Peach Mango]
	// fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 231
	highScores[2] = 143
	highScores[3] = 202                       //default allocation
	highScores = append(highScores, 392, 111) //append will reallocate the memory and all values are accomodated
	//even though defSize was 4 still NO ERROR

	fmt.Println(highScores)
	sort.Ints(highScores)
	if sort.IntsAreSorted(highScores) == true {
		fmt.Printf("Already Sorted!")
	} else {
		sort.Ints(highScores)
		fmt.Println("We sorted for you :", highScores)
	}

	
	//Removing a value from slices baseed on index.
	var list2 = []string{"javascript", "ruby", "python", "c++", "java"}
	fmt.Println(list2)
	var indToRemove int = 1
	// list2 = append(list2[indToRemove : indToRemove+1]) this will give the removed ones i.e  [ruby]
	list2=append(list2[:indToRemove],list2[indToRemove+1:]...)
	// list2=append(list2[:indToRemove],list2...) -> list2... will be init list2 fully. OUTPUT: [javascript javascript ruby python c++ java]
	fmt.Println("After removing :", list2)
}
