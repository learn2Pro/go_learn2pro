package main

import "fmt"

func main() {
	num := 11
	if num%2 == 0 {
		fmt.Println("The number ", num, " is even!")
	} else {
		fmt.Println("The number ", num, " is odd!")
	}
	num1 := 99
	if num1 <= 50 {
		fmt.Println(num1, " is less than or equal to 50")
	} else if num1 >= 51 && num1 <= 100 {
		fmt.Println(num1, " is between 51 and 100")
	} else {
		fmt.Println(num1, " is greater than 100")
	}
	if num2 := 10; num2%2 == 0 {
		fmt.Println("The number ", num2, " is even!")
	} else {
		fmt.Println("The number", num2, " is odd!")
	}
	for i := 1; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("number is %d\n", i)
	}
	fmt.Printf("break the loop")

	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("number is %d\n", i)
	}
	fmt.Printf("break the loop")

	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}
