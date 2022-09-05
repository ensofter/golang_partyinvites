package main

import "fmt"

func main() {
	counter := 0
target:
	fmt.Println("This is my target")
	counter++
	if counter < 5 {
		goto target
	}
}
