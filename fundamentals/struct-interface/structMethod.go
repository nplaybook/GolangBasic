// struct can be used as function's parameter and struct may have function-like as its field data type

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// similar to python class method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, ", my name is", customer.Name)
}

func main() {
	nplaybook := Customer{
		Name:    "nplaybook",
		Address: "On the internet",
		Age:     2,
	}
	nplaybook.sayHello("bro")
}
