package main

import "fmt"

func main() {
	names := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	allNames := names[1:]
	someNames := make([]string, 2)

	copy(someNames, allNames)

	fmt.Println("SomeNames:", someNames)
	fmt.Println("allNames:", allNames)
}
