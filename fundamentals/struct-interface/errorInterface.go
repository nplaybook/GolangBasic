// Go has a contract to make an error which using error built-in interface (nil value representing no error)

package main

import (
	"errors" // using errors built-in module
	"fmt"
)

// this function has 2 return value which is integer and error
// if there is no error, we will return nil (can be set to nil since it is an interface) instead
func divide(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New(("Division with zero"))
	} else {
		return value / divider, nil
	}
}

func main() {
	var errorExample error = errors.New("Ouch") // here we are using error built-in interface
	fmt.Println(errorExample.Error())

	// it's common to return an error for a function that possibly error
	result, err := divide(100, 0)
	if err == nil {
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
