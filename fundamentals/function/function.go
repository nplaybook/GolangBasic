// Go will execute function `main` by default. however, we can create a new function and call that function inside `main` block

package main

import "fmt"

// sometimes we need an input in our function and to accomplish that, we can add function parameter to the function itself
func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// function may return a value. to accomplish that, we need to use `return` statement and define the return data type
func getHello(name string) string {
	if name == "" {
		return "Hello annonymous"
	} else {
		return "Hello " + name
	}
}

// function may return multiple value, we need to define all returns data type
func getBirthOfDate() (string, int32) {
	return "nplaybook", 110696
}

// Go has a feature where we can create a variable from return value
// if all the variables have same data type, just define the data type at the end of last return variable
func getInfo() (fullName string, bod int32, pob string) {
	fullName = "nplaybook" // use regular = operator to assign a value to variable since the variable itself has been created when defining the function
	bod = 110696
	pob = "jkt"
	return // no need to declare the return variable explicitly but you can
}

func main() {
	// first approach
	sayHello("nplay", "book")

	// second approach, argument with variable
	firstName := "book"
	sayHello(firstName, "nplay")

	toPrint := getHello("")
	fmt.Println(toPrint)

	fullName, bod := getBirthOfDate()
	fmt.Println(fullName, "was born on", bod)

	fullName, bod, pob := getInfo()
	fmt.Println(fullName, "was born on", bod, "at", pob)
}
