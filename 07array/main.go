package main

import "fmt"

func main() {
	var list1 [3]string
	list1[0] = "jalandhar"
	list1[1] = "dhanbad"
	fmt.Println("Living in :", list1)
	fmt.Println("Len of list1 :", len(list1)) //even though contains only 2 elements, will return 3. (considers the empty space too.)

	var list2 = [4]string{"jalandhar", "delhi", "lucknow", "banaras"}
	fmt.Println("Cities :", list2)
}
