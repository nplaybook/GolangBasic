// unlike other programming languages, Go only has one looping approach which is for loop and knows no while loop
// for loop can be used for iterable data type such as array and map

package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Count in", counter)
		counter++
	}

	// we can add 2 statements in for loop which are init statement and post statement
	// init statement is a statement when the program has not been executed
	// post statement is a statement when the program will be executed once looping success

	for counter := 1; counter <= 10; counter++ { // structure: init statement; condition; post statement
		fmt.Println("Count in", counter)
	}

	// we can use for range if iteration is done for every data collection such as array, slice, or map
	slice := []string{"nplay", "book", "nplaybook"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// OR

	for index, name := range slice { // replace `index` with `_` if you dont want to use the index
		fmt.Println("Index", index, "=", name) // seems like range return 2 values which are the index and the value
	}
}
