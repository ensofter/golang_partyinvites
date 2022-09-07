package main

import "fmt"

func main() {
	names := []string{"Kayak", "Lifejacket", "Paddle"}

	appendNames := append(names, "Hat", "Gloves")

	names[0] = "Canoe"

	fmt.Println("names:", names)
	fmt.Println("appendNames:", appendNames)

	names1 := make([]string, 3, 6)

	names1[0] = "Kayak"
	names1[1] = "Lifejacket"
	names1[2] = "Paddle"

	appendNames1 := append(names, "Hat", "Gloves")

	names1[0] = "Canoe"

	fmt.Println("names1:", names1)
	fmt.Println("appendNames1:", appendNames1)

}
