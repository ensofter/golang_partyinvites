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

	valInt1 := "100"
	int1, int1err := strconv.ParseInt(valInt1, 0, 8)
	if int1err == nil {
		fmt.Println("Parsed value:", int1)
	} else {
		fmt.Println("Cannot parse value:", valInt1)
	}

	val100 := "100"

	int100, int100err := strconv.ParseInt(val100, 10, 0)
	if int100err == nil {
		fmt.Println("Parsed val:", int100)
	} else {
		fmt.Println("Cannot Parse", int100, int100err)
	}

	int1001, int1001err := strconv.Atoi(val100)

	if int1001err == nil {
		var intResult int = int1001
		fmt.Println("Poopi value:", intResult)
	} else {
		fmt.Println("cannot parse value", int1001, int1001err)
	}
}
