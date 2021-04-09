// unlike other programming language, Go is not object oriented programming
// thus, something like parent (i.e. super in Python) object doesn't exist in Go.
// however, we can cover this issue in Go by using blank interface
// example of method or function that implement blank interface is fmt.Println, panic, recover, etc
// blank interface mostly used when you want the parameters to be able to accept any data type and the function also return any data type

package main

import "fmt"

// example how to create blank interface
// return can be any data type
func Ups(num int) interface{} {
	if num == 1 {
		return 1
	} else if num == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var blank int = Ups(1)         // we cannot specify data type for blank interface, this will return an error
	var blank interface{} = Ups(1) // this will do
	blank := Ups(1)                // this will also do
	fmt.Println(blank)
}
