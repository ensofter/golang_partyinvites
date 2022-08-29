package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Value:", rand.Int())
	fmt.Println("Hello, GO!")
	fmt.Println(20 + 20)
	fmt.Println(30 + 20)

	const price float32 = 275.00
	const tax float32 = 27.50
	fmt.Println(price + tax)

	const quantity = 2
	fmt.Println("Total:", quantity*(price+tax))
}
