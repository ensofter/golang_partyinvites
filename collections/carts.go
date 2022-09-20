package main

import (
	"fmt"
	"sort"
)

func main() {
	products := make(map[string]float64, 10)

	products["Kayak"] = 279
	products["Lifejacket"] = 48.95

	fmt.Println("Map size:", len(products))
	fmt.Println("Price:", products["Kayak"])
	fmt.Println("Price:", products["Lifejacket"])

	items := map[string]float64{
		"Kayak":    279,
		"Lifejack": 48.95,
		"Hat":      0,
		"Dow":      3.74,
		"Package":  22.8,
	}
	fmt.Println("Map size:", len(items))
	fmt.Println("Price:", items["Kayak"])
	fmt.Println("Price:", items["Hat"])

	value, ok := items["Hat"]

	if ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}

	if value, ok := items["Borsh"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}

	delete(items, "Hat")

	if value, ok := items["Hat"]; ok {
		fmt.Println("Yes", value)
	} else {
		fmt.Println("No it's deleted")
	}

	for key, value := range items {
		fmt.Println("key:", key, "value:", value)
	}

	// сортировка карты
	keys := make([]string, 0, len(items))
	for key, _ := range items {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key:", key, "Value:", items[key])
	}
}
