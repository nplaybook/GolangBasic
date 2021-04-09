// in other programmning languages, if we create an object that the value hasn't been initialized then the value is null or nil
// in Go, it will be assigned to default value. Int default value is 0, string default value is "", and so on
// but, Go has nil. nil can only be used by interface, function, map, slice, and channel

package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// initialize empty map with nil
	var person map[string]string = nil
	// we can use nil to check data availability
	if person == nil {
		fmt.Println("Empty map")
	} else {
		fmt.Println(person)
	}

	// if we don't use nil, the script will be less efficient
	var nextPerson map[string]string
	if nextPerson["name"] == "" {
		fmt.Println("Empty data")
	}

	eko := newMap("eko")
	fmt.Println(eko)
}
