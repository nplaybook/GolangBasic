// there are 2 numeric data type: Floating, Integer
// in Go, there are more of it (depends on capacity)

// ::int8 (-128 - 127)
// ::int16 (-32768 - 32767)
// ::int32 (-214783648 - 214783647) <alias> rune / int
// ::int64 (-9223372036854775808 - 9223372036854775807) <alias> int

// ::uint8 (0 - 255) <alias> byte
// ::uint16 (0 - 65636)
// ::uint32 (0 - 4294967295) <alias> uint
// ::uint64 (0 - ....................) <alias> uint

// ::float32 (1.18 x ^-38 - 3.4 x ^38)
// ::float64 (2.23 x ^-308 - 1.8 x ^308)
// ::complex64
// ::complex128 

package main

import "fmt"

func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Due = ", 2)
	fmt.Println("TigaKomaLima = ", 3.5)
}