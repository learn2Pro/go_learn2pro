package main

import (
	"fmt"
	"math/rand"
)

func main() {
	finger := 4
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("incorrect finger number")
	}

	switch num := 25; {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
	}

	switch num1 := -1; {
	case num1 < 50:
		fmt.Printf("%d is less than 50\n", num1)
		if num1 < 0 {
			break
		}
		fallthrough
	case num1 < 100:
		fmt.Printf("%d is less than 100\n", num1)
		fallthrough
	case num1 < 200:
		fmt.Printf("%d is less than 200\n", num1)

	}
randloop:
	for {
		switch i := rand.Intn(100); {
		case i%2 == 0:
			{
				fmt.Printf("Generated num %d is even", i)
				break randloop
			}
		}
	}

}
