package main

import "fmt"

func printPrice() {
	kayakPrice := 275.00
	kayakTax := kayakPrice * 0.2
	fmt.Println("Price:", kayakPrice, "Tax:", kayakTax)
}

func printPrice1(product string, price float64, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

func printSuppliers(product string, suppliers []string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

func printSuppliers1(product string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Println("Product1:", product, "Supplier:", supplier)
	}
}

func printSuppliers2(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
	}
}

func swapValues(first, second *int) {
	fmt.Println("Before swap:", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("After swap:", *first, *second)
}

func main() {
	fmt.Println("About to call function")
	printPrice()
	fmt.Println("Function complite")

	printPrice1("Kayak", 275, 0.2)
	printPrice1("Lifejacket", 48.95, 0.2)
	printPrice1("Soccer Ball", 19.50, 0.15)

	printSuppliers("Kayak", []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"})
	printSuppliers("Lifejacket", []string{"Sail Safe Co"})
	fmt.Println("________________________________________________________")
	printSuppliers1("Kook", "Acme Kayaks", "BObs Boats", "Crazzy Canoe")
	printSuppliers1("Lifejacket", "Safe COmpanny CO")
	fmt.Println("________________________________________________________")
	printSuppliers2("Kayak", "Acme Kayaks", "Bob;s Boats", "Crazy Canoes")
	printSuppliers2("LafiJacket", "Sail Safe Co")
	printSuppliers2("Soccer Ball")
	fmt.Println("________________________________________________________")
	names := []string{"Puting Hiy", "Posrat na hohlov", "Ebal v rot mobilizaciu"}
	printSuppliers2("Kayak", names...)
	fmt.Println("________________________________________________________")
	val1, val2 := 10, 20
	fmt.Println("Before calling function", val1, val2)
	swapValues(&val1, &val2)
	fmt.Println("After calling function", val1, val2)
}
