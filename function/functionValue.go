// in Go, function is also a data type where it can be assigned to a variable
// this approach works good for if we have a function that requires another function as its paramter
// this approach supports our modular function

package main

import "fmt"

func getGoodBye(name string, filter func(string) string) string { // filter data type is a function with string as its parameter and return string
	return "Good Bye " + filter(name)
}

// in getGoodBye filter parameter, we can there might be an issue if our function has a lot of parameters and we need to define all the data types
// to handle that issue, we can use type declarations
type Filter func(string) string

func sayHello(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func nameFilter(name string) string {
	if name == "nplaybook" {
		return name
	} else {
		return "***"
	}
}

func main() {
	filter := nameFilter
	fmt.Println(getGoodBye("nplay", filter))
	fmt.Println(getGoodBye("nplaybook", nameFilter)) // use the function name directly

	sayHello("nplaybook", nameFilter)
}
