// the last parameter in a function has a capability to be a `varargs`
// `varargs` means a parameter might accepts more than 1 input or an Array or a Slice
// a function only has 1 varargs

package main

import "fmt"

func sumAll(numbers ...int) int { // varargs has `...`
	total := 0
	for _, number := range numbers { // similar to enumerate in Python
		total += number
	}
	return total
}

func main() {
	empty := sumAll() // it's okay not to input any
	total := sumAll(10, 10, 10, 5)
	fmt.Println(empty)
	fmt.Println(total)

	// input a Slice in varargs
	sliceNumbers := []int{10, 10, 10, 10}
	sliceTotal := sumAll(sliceNumbers...)
	fmt.Println(sliceTotal)
}
