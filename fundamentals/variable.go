// one variable only has one data type
// use `var` to create a variable
// a variable should be used in the program
// if a variable is not used, the program returns error

package main

import "fmt"

func main(){
	var name string // declare data type
	
	name = "nplaybook"
	fmt.Println(name)

	// update with new string
	name = "booknplay"
	fmt.Println(name)

	// Go also accept a variable with no data type declaration
	// only if the variable is given a value. Go will automatically
	// identify the data type
	var firstName = "nplay"
	fmt.Println(firstName)

	// if you assign a non decimal number, Go will identify
	// as int32/64 unless you specify the int data type
	var age = 18
	var nextDate int8 = 1
	fmt.Println(age)
	fmt.Println(nextDate)

	// we can use := operator to replace `var` 
	// but we need to give the value right away
	// this way we can't declare data type
	height := 168
	fmt.Println(height)

	// declare multiple variable
	var (
		lastName = "book"
		weight float32 = 65.5
	)
	fmt.Println(lastName)
	fmt.Println(weight)
}