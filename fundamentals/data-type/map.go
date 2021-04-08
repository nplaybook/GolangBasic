// in array or slice, we use index to access the data
// map is another data type that consist of collection of items. however, in map we may define data type of the index, we may add
// as many items as we want in key-value pair (the key has to be unique otherwise older data with same key will be replaced)

package main

import "fmt"

func main() {
	person := map[string]string{
		"name":     "nplaybook",
		"language": "golang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["language"])

	// add more data
	person["job"] = "explorer"
	fmt.Println(person)

	// functions in map
	// len(map) -> get the length or total data
	// map[key] -> get the value from a specific key
	// map[key] = value -> create new key-value pair or replace a value from specific key
	// make(map[Typekey]TypeValue) -> create a new map
	// delete(map, key) -> delete data or value from a specific key

	// create empty map
	team := make(map[string]string)
	team["captain"] = "LeBron"
	team["playmaker"] = "Schroeder"
	team["bench"] = "Kuz"
	fmt.Println((team))

	// delete a key-value
	delete(team, "bench")
	fmt.Println((team))
}
