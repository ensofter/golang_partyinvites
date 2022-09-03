package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1, val2, val3 := "true", "false", "true"

	bool1, b1err := strconv.ParseBool(val1)
	bool2, b2err := strconv.ParseBool(val2)
	if bool3, b3err := strconv.ParseBool(val3); b3err != nil {
		panic(b3err)
	} else {
		fmt.Println("Bool 3", bool3, b3err)

	}

	fmt.Println("Bool 1", bool1, b1err)
	fmt.Println("Bool 2", bool2, b2err)
}
