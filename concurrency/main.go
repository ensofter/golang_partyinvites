package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main function started")
	CalcStoreTotal(Products)
	time.Sleep(time.Second * 5)
	fmt.Println("main function complete")
}
