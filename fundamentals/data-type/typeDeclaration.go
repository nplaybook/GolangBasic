// we can create a new data type by using type declaration
// type delcaration mosly used for alias

package main

import "fmt"

func main() {
	type NoID string

	var id NoID = "123456789"
	fmt.Println(id)
}
