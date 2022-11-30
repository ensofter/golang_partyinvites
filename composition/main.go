package main

import (
	"composition/store"
	"fmt"
)

func main() {
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}
	for _, b := range boats {
		fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Name, "Price:", b.Price(0.2))
	}
}
