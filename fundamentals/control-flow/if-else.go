package main

import "fmt"

func main() {
	name := "nplaybook"

	if name == "book" {
		fmt.Println(name)
	} else if name == "nplaybook" {
		fmt.Println("Match")
	} else {
		fmt.Println("Not match")
	}

	// if short statement
	if length := len(name); length > 2 {
		fmt.Println(true)
	} else {
		fmt.Println((false))
	}
}
