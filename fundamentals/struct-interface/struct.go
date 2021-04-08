// struct is template data (representation) that is used to merge 0 or more other data type into one
// simply put, struct is a collection of fields
// the issue comes with array data type that may only store 1 data type

package main

import "fmt"

// it's like making a class in OOP
// we need to define struct first before use it and we can't use it direcly while creating it
// notes: not necessarily has to fill every fields
type Customer struct {
	Name, Address string // similar to python class attribute
	Age           int
	Hobby         string
}

func main() {
	// create variable with struct data type
	var nplaybook Customer
	nplaybook.Name = "nplaybook"
	nplaybook.Address = "On the internet"
	nplaybook.Age = 2

	fmt.Println(nplaybook)
	fmt.Println(nplaybook.Age)

	// another way to create struct by struct literals
	nplay := Customer{
		Name:    "nplay",
		Address: "None",
		Age:     0,
		Hobby:   "Typing",
	}
	fmt.Println(nplay)
}
