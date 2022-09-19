package main

import "fmt"

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
}
