// interface is a abstract data type, doesn't have direct implementation
// interface contains of method's definitions and commonly used as a contract
// interface can be used as function's parameter

// overview: whatever the data type that has function/method similar to interface,
// it can use any funtion which asssigning that interface as its parameter
// a bit similar to python parent and inherent class

package main

import "fmt"

type HasName interface {
	GetName() string // method GetName returns string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

// since Person data type has GetName method which makes it automatically follow HasName contract,
// Person data type may call sayHello function
func (friend Person) GetName() string {
	return friend.Name
}

func main() {
	friend := Person{Name: "booker"}
	sayHello(friend)
}
