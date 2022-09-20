package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello")

	var price string = "$48.95"

	var currency string = string(price[0])
	var amountString string = price[1:]

	amount, parseErr := strconv.ParseFloat(amountString, 64)

	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse erroe:", parseErr)
	}

	var price1 []rune = []rune("€48.95")

	var currency1 string = string(price1[0])
	var amountString1 string = string(price1[1:])
	amount1, parseError1 := strconv.ParseFloat(amountString1, 64)

	fmt.Println("Length:", len(price1))
	fmt.Println("Currency:", currency1)
	if parseError1 == nil {
		fmt.Println("Amount:", amount1)
	} else {
		fmt.Println("Erroe is occured", parseError1)
	}

	for index, char := range price {
		fmt.Println(index, char, string(char))
	}
}
