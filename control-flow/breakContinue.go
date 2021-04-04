// break and continue is used in for loop
// use break if you want to stop the for loop program
// use continue if you want to skip one for loop iteration and do the next iteration

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Println("Itteration number", i)
	}
}
