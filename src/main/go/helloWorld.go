package main

import (
	"fmt"
)

func main() {
	var age int = 29 // 声明变量并初始化
	fmt.Println("my age is", age)
	age = 54
	fmt.Println("my age is ", age)
	age = 108
	fmt.Println("my age is ", age)
	fmt.Println("Hello World")
	var width, height int = 100, 50 // 声明多个变量
	fmt.Println("width is:", width, "height is:", height)
	var (
		name    = "naveen"
		age1    = 29
		height1 int
	)
	fmt.Println("my name is: ", name, ", age is", age1, "and height is", height1)
	name2, age2 := "abc", 103
	fmt.Println("my name is:", name2, ",age is", age2)

	//a, b := 145.8, 543.8
	//c := math.Min(a, b)
	//fmt.Println("minimum value is ", c)

	//var a int = 89
	//b := 95
	//fmt.Println("value of a is", a, "and b is", b)
	//fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   // a 的类型和大小
	//fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小

	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b%T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum:", sum, "diff:", diff)

	no1, no2 := 56, 89
	fmt.Println("sum:", no1+no2, "diff", no1-no2)

	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	i := 55            //int
	j := 67.8          //float64
	sum0 := i + int(j) //j is converted to int
	fmt.Println(sum0)

	intI := 10
	var floatJ float64 = float64(intI) // 若没有显式转换，该语句会报错
	fmt.Println("floatJ:", floatJ)
}
