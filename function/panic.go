// panic function is a function that we can use to stop a program
// panic function mostly used if error occurs in our program
// when panic function calls, program will stop but defer function will be executed anyway

package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	} else {
		fmt.Println("App running")
	}
}

func main() {
	runApp(false)
	runApp(true)
}
