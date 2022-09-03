package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := true
	val2 := false

	str1 := strconv.FormatBool(val1)
	str2 := strconv.FormatBool(val2)

	fmt.Println("Formatted value 1:" + str1)
	fmt.Println("Formatted value 2:" + str2)

	val3 := 275

	base10String := strconv.FormatInt(int64(val3), 10)
	base2String := strconv.FormatInt(int64(val3), 2)

	fmt.Println("Base 10: " + base10String)
	fmt.Println("Base 2: " + base2String)
}
