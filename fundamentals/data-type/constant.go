// constant is imutable variable
// use `const` to declare
// must give the value right away
// Go will not complain if a constant is not used (unlike variable)

package main

import "fmt"

func main() {
	const firstName string = "nplay" // don't have to declare data type
	const lastName = "book"

	// error
	firstName = "book"
	lastName = "nplay"
	fmt.Println(firstName)
	fmt.Println(lastName)

	// multiple declartion
	const (
		firstName string = "nplay"
		lastName         = "book"
	)
}
