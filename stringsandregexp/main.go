package main

import (
	"fmt"
	"strings"
	"unicode"
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

	product1 := "Kayak"
	for _, char := range product1 {
		fmt.Println(string(char), "Upper case:",
			unicode.IsUpper(char))
	}

	description1 := "A boat for one person"
	fmt.Println("Count:", strings.Count(description1, "o"))
	fmt.Println("Index:", strings.Index(description1, "o"))
	fmt.Println("LastIndex:", strings.LastIndex(description1,
		"o"))
	fmt.Println("IndexAny:", strings.IndexAny(description1,
		"abcd"))
	fmt.Println("LastIndex:", strings.LastIndex(description1,
		"o"))
	fmt.Println("LastIndexAny:",
		strings.LastIndexAny(description1, "abcd"))

	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	fmt.Println("IndexFunc:", strings.IndexFunc(description1, isLetterB))
}
