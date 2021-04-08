// recursive function is a function that calls itself
// one of the example of recustive function implementation is factorial

package main

import "fmt"

// factorial solution without implement recursive
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// implement recursive
func factorialRecursive(value int) int {
	if value == 1 { // this is the base case or stopping scenario
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(5))
	fmt.Println(factorialRecursive(5))
}
