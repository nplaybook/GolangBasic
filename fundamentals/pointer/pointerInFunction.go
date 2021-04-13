// when we create a parameter in a function, by default the approach is `pass by value`
// so, whenever we change the data in a function, the original will remain the same
// but, sometimes we want to make a change for a value that is passed onto a function
// to accomplish that, we can use operator *

// consider to use pointer when we have a big struct data and calls multiple time within functions
// it may harm the memory

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeAddress(address Address) {
	address.Country = "Indonesia"
}

// notice the difference on * operator
func changeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{
		City:     "South Tangerang",
		Province: "Banten",
		Country:  "",
	}

	changeAddress(address) // no data change
	fmt.Println(address)
	changeAddressToIndonesia(&address) // changes here, notice the difference in & operator
	fmt.Println(address)
}
