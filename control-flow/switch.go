// switch expression can be used to control which program to run (just like if expression)
// most of the time, switch statement is used to check condition of a variable

package main

import "fmt"

func main() {
	name := "nplaybook"
	switch name { // switch statement to initiate which variable we need to check
	case "nplaybook":
		fmt.Println("Hello nplaybook")
	case "book":
		fmt.Println("Hello book")
	default: // similar to else in if statement, if every case statement fails then run default
		fmt.Println("What is your name?")
	}

	// switch with short statement
	// no need to use default since in boolean check there are only true or false
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name is too long")
	case false:
		fmt.Println("Okay")
	}

	// switch with no condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name is too long")
	case length < 5:
		fmt.Println("Name is too short")
	default:
		fmt.Println("Okay")
	}
}
