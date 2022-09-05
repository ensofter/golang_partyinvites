package main

import "fmt"

func main() {
	product := "Kayakuuuu"
	for index, char := range product {
		switch char {
		case 'K', 'k':
			if char == 'k' {
				fmt.Println("Lower k at position", index)
				break
			}
			fmt.Println("Upper K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		default:
			fmt.Println("7uuuuu")
		}
	}

}
