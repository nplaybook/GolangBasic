// use to find the result from 2 boolean values
// && -> and
// || -> or
// ! -> opposite (!false == true)

package main

import "fmt"

func main() {
	var (
		finalScore     int8 = 90
		presence       int8 = 60
		passFinalScore bool = finalScore > 80
		passPresence   bool = presence > 75
		pass           bool = passFinalScore && passPresence
	)

	fmt.Println(pass)

	fmt.Println(true && false)
	fmt.Println(false || true)
	fmt.Println(!false == true)
}
