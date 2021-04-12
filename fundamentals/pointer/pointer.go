// by default, all variables in Go is `passing by value`, not `by reference`
// that means if we pass a variable to function's argument, method or other variables, what happen is a value from variable will be duplicated and pass on to those
// memory-wise, original and duplicate value will be saved in different location

// with Pointer, we can make a reference to same memory location without having to duplicate these value
// put it simply, we are now using pass by reference instead of pass by value
// to make a pointer in Go, we can use operator `&` followed by name of the variable

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "South Tangerang",
		Province: "Banten",
		Country:  "Indonesia",
	}
	address2 := address1 // assign address1 to address2 to check if we modify item in address2, whether item in address1 is modified or not

	address2.City = "Tangerang"

	// no changes in address1 since the value in address1 is duplicated and assigned to address2. no direct relationship between address1 and address2
	// in other programming language such as C or java it will be changed (using by reference approach)
	fmt.Println(address1)
	fmt.Println(address2)

	// address3 use pointer to address1, pass by reference
	address3 := &address1 // OR var address3 *Address = &address1

	address3.City = "Serang"

	// address1 City is now Jakarta
	fmt.Println(address1)
	fmt.Println(address3)

	// above, we change a value for a field
	// now, we want to change the whole value but keep the referenced value the same
	address3 = &Address{
		City:     "Bogor",
		Province: "West Java",
		Country:  "Indonesia",
	}

	// in this case, address1 is {Serang, Bantent, Indonesia} meanwhile address3 is {Bogor, West Java, Indonesia}
	// notice that address1 still receive an impact from City value modification but not following latest address3 modification (pointed to another reference)
	fmt.Println(address1)
	fmt.Println(address3)

	// we can also change address1 (previously as pointed variable) to new memory
	// to accomplish that, we need to use `*` operator
	// nb: delete address3 = &Address{} to see changes
	*address3 = Address{
		City:     "Pekanbaru",
		Province: "Riau",
		Country:  "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address3)

	// Go has built-in function `new` to create a pointer
	// however, tihs function returns empty data
	newAddress := new(Address)
	newAddress2 := newAddress

	newAddress2.Country = "German"

	fmt.Println(newAddress) // initiated by no data but due to changes now they have German in Country
	fmt.Println(newAddress2)
}
