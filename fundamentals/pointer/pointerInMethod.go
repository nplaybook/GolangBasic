// eventho method attaches to struct, in fact struct that is accessed in a method is `pass by value`
// highly recommended to use pointer in method to avoid memory issue

package main

import "fmt"

type Man struct {
	Name string
}

type Woman struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

// notice the difference on * operator
func (man *Man) Master() {
	man.Name = "Guru " + man.Name
}

func main() {
	nplaybook := Man{Name: "nplaybook"}
	nplaybook.Married()         // since Married method not using pointer, nplaybook data will be duplicated and sent to Married method
	fmt.Println(nplaybook.Name) // no changes, not using pointer

	nplaybook.Master()
	fmt.Println(nplaybook.Name)
}
