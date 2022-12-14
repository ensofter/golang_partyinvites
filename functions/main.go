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

func calcTax(price float64) float64 {
	return price + (price * 0.2)
}

func swapValues1(first, second int) (int, int) {
	return second, first
}

func calcTax1(price float64) float64 {
	if price > 100 {
		return price * 0.2
	}
	return -1
}

func calcTax2(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcTax2(price); due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	return
}

func calcTotalPrice1(products map[string]float64) (count int, total float64) {
	count = len(products)
	for _, price := range products {
		total += price
	}
	return
}

func calcTotalPrice2(products map[string]float64) (count int, total float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer start")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Secend defer start")
	fmt.Println("Function about to return")
	return
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
	fmt.Println("________________________________________________________")
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products {
		priceWithTax := calcTax(price)
		fmt.Println("Product:", product, "Price:", priceWithTax)
	}
	fmt.Println("________________________________________________________")
	val3, val4 := 10, 20
	fmt.Println("Before calling function", val3, val4)
	val3, val4 = swapValues1(val3, val4)
	fmt.Println("After calling function", val3, val4)
	fmt.Println("________________________________________________________")

	products1 := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products1 {
		tax := calcTax1(price)
		if tax != -1 {
			fmt.Println("Product:", product, "Tax", tax)
		} else {
			fmt.Println("Product:", product, "No tax due")
		}
	}
	fmt.Println("________________________________________________________")
	total1, tax1 := calcTotalPrice(products1, 10)
	fmt.Println("Total 1:", total1, "Tax 1:", tax1)
	fmt.Println("________________________________________________________")
	_, total := calcTotalPrice1(products)
	fmt.Println("Total:", total)
	fmt.Println("________________________________________________________")
	z, total1 := calcTotalPrice2(products)
	fmt.Println("Total:", total1)
	fmt.Println(z)
}
