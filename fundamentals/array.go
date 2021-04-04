// array consists of collection of items that have same data type
// when we create an array, we need to define how many data that array may contain
// once you define max number of data in an array, you cannot add more data than the max number

package main

import "fmt"

func main() {
	// first approach
	var names [3]string
	names[0] = "nplay"
	names[1] = "book"
	names[2] = names[0] + names[1]
	// names[3] = "try" // error will appear telling you that the index is out of bound since we define max number of variables names is 3 items

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// second approach
	var values = [3]int{
		75, 65, 50, // last item should have a , as well
	}
	fmt.Println(values)

	// there are other functions we can use for array
	// len(array) -> get the length (max number) of an array not number of data has been registered
	// array[index] -> get item based on index position
	// arrau[index] -> change item in a specific index
	values[0] = 10
	fmt.Println(values)
	fmt.Println(len(values))
}
