// closure is a function's ability to interact with others data within the same scope
// the scope itself can be determined by function block `{}`

package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		fmt.Println("Incerement")
		counter++           // counter variable is accessible since increment function is within the main block
		name := "nplaybook" // this variable is not accessible outside increment function block
		fmt.Println(name)
	}

	increment()
	fmt.Println(counter)
	fmt.Println(name) // this will return an error
}
