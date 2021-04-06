// most of the times we create a function using `func` but there is another way to create a function
// we can assign it to a variable or to a parameter

package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// #1 assign function to a variable
	blacklist := func(name string) bool {
		return name == "nplaybook"
	}

	registerUser("nplaybook", blacklist)

	// #2 create a function in paramaeter
	registerUser("nplay", func(name string) bool {
		return name == "nplaybook"
	})
}
