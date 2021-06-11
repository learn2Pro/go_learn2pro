package main

import (
	"fmt"
	"log"
)

var p, r, t = 5000.0, 10.0, 1.0

func init() {
	println("Main Package initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

//Calculate calculates and returns the simple interest for a principal p, rate of interest r for time duration t years
func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
func main() {
	fmt.Println("Simple interest calculation")
	si := Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
