// slice is a part of an array which access subset or the whole array
// slice has 3 components:
// 1. pointer -> refers to the starting index of slice from an array
// 2. length -> length of the slice
// 3. capacity -> slice's capatiy (max length) which starting from the pointer index and last index of an array

package main

import "fmt"

func main() {
	// use ... in array if you are not sure how many data this array may contain
	// if we don't specify specifically or use ... (means we are letting the [] empty), it will be treated as a slice, be carefull!
	var months = [...]string{
		"January",
		"Febrauary",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	// pointer -> 6
	// length -> 3 (last index - start index)
	// capacity -> 6 (max number of items of an array - start index)
	var holiday = months[6:9]
	fmt.Println(holiday)

	// when we modifiy data in slice, the array data is also modified
	holiday[0] = "birthdayMonth"
	fmt.Println(holiday)
	fmt.Println(months)

	// there are some functions we can use for slice
	// len(slice) -> get the length of a slice
	// cap(slice) -> get the capacity of a slice
	// append(slice, data) -> create a new slice by add new data in the last index. if the capacity is full, will create new array
	// make([]TypeData, length, capacity) -> create a new slice
	// copy(destination, source) -> copy paste a slice from a source to a destination
	fmt.Println(len(holiday))
	fmt.Println(cap(holiday))
	var makeHoliday = append(holiday, "newMonth")
	fmt.Println(makeHoliday)
	fmt.Println(holiday)
	// October will be replaced by newMonth since variable holiday still has some capacity to be added by new item and variable months is full
	fmt.Println(months)

	// safer way to create a slice
	var newSlice = make([]string, 2, 5)
	newSlice[0] = "nplay"
	newSlice[1] = "book"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice from a variable to another variable
	fromSlice := months[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // make sure the length and capacity are not cut
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
}
