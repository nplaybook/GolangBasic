/**
recover is a function that we can use to catch panic
by using recover, panic will be stopped and program will keep running
*/

package main

import "fmt"

func endApp() {
	fmt.Println("Program end")
	// best practice to implement recover function is by write it within defer function
	message := recover() // string error within panic function will be catch by recover
	if message != nil {
		fmt.Println("Error occured:", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Somehow error")
	}
	fmt.Println("Program running")
}

func main() {
	runApp(true)
	fmt.Println("Program keep running") // this program will be executed even an error occured
}
