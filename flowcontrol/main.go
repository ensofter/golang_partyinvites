package main

import (
	"fmt"
)

func main() {
	product := "koyak"

	for index, character := range product {
		fmt.Println("Index:", index)
		fmt.Println("Char:", string(character))
	}
}
