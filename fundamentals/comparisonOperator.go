// use to compare 2 values with boolean data type as the result
// you can use <. >, ==, !=, <=, and >=

package main

import "fmt"

func main() {
	var (
		firstName = "nplay"
		lastName  = "book"
	)
	fmt.Println(firstName == lastName)
	fmt.Println(firstName != lastName)

	lastName = "nplay"
	fmt.Println(firstName == lastName)
	fmt.Println(firstName > lastName) // will compare based on sum of unicode
}
