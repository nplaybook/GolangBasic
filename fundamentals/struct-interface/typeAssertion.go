// with type assertion we can change a data type into another data type
// this feature mostly used when we are dealing with blank interface

package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	// resultInt := result.(int) // panic due to string is forcibly convert to int, we can use switch to avoid error
	// data type change in case expression
	switch diffResult := result.(type) {
	case string:
		fmt.Println("String", diffResult)
	case int:
		fmt.Println("Int", diffResult)
	default:
		fmt.Println("Unknown")
	}
}
