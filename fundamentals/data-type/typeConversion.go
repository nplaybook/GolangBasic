// convert data type within one domain: int32 to int64, int32 to float32
// please be carefull when you try to convert int64 to int16
// as int16 might be not enough to store the value
// there will be no error but the value will change

package main

import "fmt"

func main() {
	var (
		num32 int32 = 10000
		num64 int64 = int64(num32) // convert
		num8  int8  = int8(num32)
	)
	fmt.Println(num32)
	fmt.Println(num64)
	fmt.Println(num8)

	var (
		name       = "nplaybook"
		firstIndex = name[0]
		firstChar  = string(firstIndex) // char ASCII (in byte) to string
	)
	fmt.Println(firstIndex)
	fmt.Println(firstChar)
}
