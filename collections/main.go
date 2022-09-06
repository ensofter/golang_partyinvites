package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"

	fmt.Println(names)

	names1 := [3]string{"Kayak", "Lifejacket", "Paddle"}

	fmt.Println(names1)

	otherArray := &names

	names[0] = "Penis"

	fmt.Println("names", names)
	fmt.Println("otherArray", *otherArray)

	names[0] = "Kayak"
	same := names == names1
	fmt.Println("comparison:", same)

	names2 := [4]string{"Kayak", "Lifejacket", "Paddle"}
	for index, value := range names2 {
		fmt.Println("Index:", index, "Value:", value)
	}

	names3 := make([]string, 3)

	names3[0] = "Kayak"
	names3[1] = "Lifejacket"
	names3[2] = "Paddle"

	fmt.Println(names3)

	names4 := []string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names4)

	names4 = append(names4, "Hat", "Gloves")

	fmt.Println(names4)

	names5 := make([]string, 3, 6)
	names5[0] = "Kayak"
	names5[1] = "Lifejacket"
	names5[2] = "Paddle"

	fmt.Println("Len:", len(names5))
	fmt.Println("Cap:", cap(names5))

	moreNames := []string{"Hat Gloves"}
	appendedNames := append(names5, moreNames...)

	fmt.Println("appendedNames:", appendedNames)
}
