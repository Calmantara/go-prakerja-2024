package main

import (
	"fmt"
)

type Status = string

const (
	STATUS_ONLINE Status = "ONLINE"
)

// standard minimum untuk menjalankan
// program golang
// 1. pakcage main
// 2. func main
func main() {
	// I. declare variable
	// 1. with data type
	var name string = "calman"
	var age uint
	age = 17
	fmt.Println(name, age)
	// 2. without data type
	email := "calman@mail.com"
	fmt.Println(email)

	fmt.Print("this is print", email, "\n")
	fmt.Printf("data type for email variable:%T %f %s\n", email, 3.14, name)

	// II. data type
	var address string
	var isValid bool
	var cost float64

	var ui uint8
	var b byte

	var addressPtr *string
	fmt.Println(address, isValid, cost, b, ui, addressPtr)

	// nil => null => nihil
	// variable yang tidak ada alamat / memory
	// nil == ""
	// nil == 0

	zeroString := "0"
	zeroInt := 0
	// fmt.Println(zeroString == zeroInt)
	// 1 + "1" = "11"
	fmt.Println(zeroString+"calman", zeroInt)

	// string assign
	name = "'calman'"
	fmt.Println(name)
	name = `"'calman'"`
	fmt.Println(name)

	// interface
	var varInterface interface{}
	varInterface = 1
	fmt.Println(varInterface)
	varInterface = "calman"
	fmt.Println(varInterface)
}
