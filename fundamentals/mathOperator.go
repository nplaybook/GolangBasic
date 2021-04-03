// you can use +, -, * and / for basic math operation
// OR you can use augmented assignments if you want to do math operation on the same variable
// OR you can use unary operator: ++ equal to a = a + 1, -- equal to a = a - 1, - equal to negative, + equal to positive, ! equal to reverse

package main

import "fmt"

func main() {
	var (
		num1   int8 = 5
		num2   int8 = 10
		result      = num1 + num2
	)
	fmt.Println(result)

	var num3 int8 = 25
	num3 += 5
	fmt.Println(num3)

	num3++
	fmt.Println(num3)

}
