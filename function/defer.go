// defer function is a function that we can schedule to execute after another function has been executed
// defer function will always be executed even the previous function return an error

package main

import "fmt"

func logging() {
	fmt.Println("Done executing function")
}

func runApplication(value int) {
	defer logging() // this function will be executed if fmt.Println("Run Application") OR runApplication function is executed
	// defer function has to be put on the top of the function, not on the bottom of the function. otherwise, if error appear it won't be executed
	// error message will appear anyway
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(0)
}
