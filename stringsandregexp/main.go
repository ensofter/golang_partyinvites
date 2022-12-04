package main

import (
	"fmt"
	"strings"
)

func main() {
	product := "Kayak"

	fmt.Println("Product:", product)

	fmt.Println("Contains:", strings.Contains(product,
		"yak"))
	fmt.Println("ContainsAny:", strings.ContainsAny(product,
		"abc"))
	fmt.Println("ContainsRune:",
		strings.ContainsRune(product, 'K'))
	fmt.Println("EqualFold:", strings.EqualFold(product,
		"KAYAK"))
	fmt.Println("HasPrefix:", strings.HasPrefix(product,
		"Ka"))
	fmt.Println("HasSuffix:", strings.HasSuffix(product,
		"yak"))

	description := "привет как твои дела"
	fmt.Println("Original:", description)
	fmt.Println("Title:", strings.ToUpper(description))
}
