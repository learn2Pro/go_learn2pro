package main

import (
	"fmt"
	"math"
	"unsafe"
)

/*
bool Types
Numeric Types
	int8, int16, int32, int64, int
	uint8, uint16, uint32, uint64, uint
	float32, float64
	complex64, complex128
	byte
	rune
string Types
*/
//int8: represents 8 bit signed integers
//size: 8 bits
//range: -128 to 127
//
//int16: represents 16 bit signed integers
//size: 16 bits
//range: -32768 to 32767
//
//int32: represents 32 bit signed integers
//size: 32 bits
//range: -2147483648 to 2147483647
//
//int64: represents 64 bit signed integers
//size: 64 bits
//range: -9223372036854775808 to 9223372036854775807
func main() {
	//short hand syntax
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	var c = a && b
	fmt.Println("c:", c)
	var d = a || b
	fmt.Println("d:", d)

	var aa int = 89
	bb := 95
	fmt.Println("value of a is", aa, "and b is", bb)
	fmt.Printf("type of a is %T, size of a is %d", aa, unsafe.Sizeof(aa))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", bb, unsafe.Sizeof(bb)) //type and size of b

	const ca = 5
	fmt.Println("const of a is:", ca)

	const (
		name    = "John"
		age     = 50
		country = "Canada"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	af, bf := 145.8, 543.8
	cf := math.Min(af, bf)
	fmt.Println("Minimum value is", cf)

	const n = "Sam"
	var nameVar = n
	fmt.Printf("type %T value %v", nameVar, nameVar)

	var aDiv = 5.9 / 8
	fmt.Printf("a's type is %T and value is %v", aDiv, aDiv)

}
func calculateBill(price, number int) int {
	return price * number
}
func calculateArea(length, width int) (int, int) {
	return length * width, 2 * (length + width)
}
