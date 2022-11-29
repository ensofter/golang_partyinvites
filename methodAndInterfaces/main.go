package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
}

func (s *Service) getName() string {
	return s.description
}

func (s *Service) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}

type Product struct {
	name, category string
	price          float64
}

func (p Product) getName() string {
	return p.name
}

func (p Product) getCost(_ bool) float64 {
	return p.price
}

type Accout struct {
	accountNumber int
	expenses      []Expense
}

func main() {
	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		&Service{"Boat Cover", 12, 89.50},
	}

	for _, expense := range expenses {
		fmt.Println("Product:", expense.getName(), "Price:", expense.getCost(true))
	}
	fmt.Println("Total:", calcTotal(expenses))

	account := Accout{
		accountNumber: 12345,
		expenses: []Expense{
			Product{"Kayak", "Watersports", 275},
			&Service{"Boat Cover", 12, 89.50},
		},
	}

	for _, expense := range account.expenses {
		fmt.Println("Expenses:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total:", calcTotal(account.expenses))

	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = &product
	product.price = 100
	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost(false))

	service := Service{"Boat Cover", 12, 89.50}
	var expense1 Expense = &service
	service.monthlyFee = 1001
	fmt.Println("Product field value:", service.monthlyFee)
	fmt.Println("Expense method result:", expense1.getCost(false))

	var e1 Expense = Product{name: "Kayak"}
	var e2 Expense = Product{name: "Kayak"}
	var e3 Expense = &Service{description: "Boat Cover"}
	var e4 Expense = &Service{description: "Boat Cover"}
	fmt.Println("e1 == e2", e1 == e2)
	fmt.Println("e3 == e4", e3 == e4)

	expenses3 := []Expense{
		Product{"Kayak", "Watersports", 275},
		Product{"Ball", "Watersports", 300},
		&Service{"Boat Cover", 12, 89.50},
	}

	for _, expense3 := range expenses3 {
		if s, ok := expense3.(Product); ok {
			fmt.Println("Product:", s.name, "Category:", s.category)
		} else {
			fmt.Println("Expense:", expense3.getName(), "Cost:", expense3.getCost(true))
		}
	}
}
