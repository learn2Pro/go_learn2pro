package main

import "fmt"

func main() {
	var n [3]int
	fmt.Println(n)

	n[0] = 12
	n[1] = 13
	n[2] = 14
	fmt.Println(n)

	var a = [3]int{1, 2, 3}

	fmt.Println(a)
	b := [3]int{1, 2, 3}
	b[0] = 100
	fmt.Println(b)

	c := [3]int{12}
	fmt.Println(c)

	d := [...]int{12, 78, 50, 100} // ... makes the compiler determine the length
	fmt.Println(d)

	aa := [...]string{"USA", "China", "India", "Germany", "France"}
	bb := aa // a copy of a is assigned to b
	bb[0] = "Singapore"
	fmt.Println("aa is ", aa)
	fmt.Println("bb is ", bb)

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)

	af := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of af is", len(af))
	for i := 0; i < len(af); i++ {
		fmt.Printf("the %d index, the value %.2f\n", i, af[i])
	}
	for i, v := range af {
		fmt.Printf("the %d index, the value %.2f\n", i, v)
	}

	multiA := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(multiA)
	var multiB [3][2]string
	multiB[0][0] = "apple"
	multiB[0][1] = "samsung"
	multiB[1][0] = "microsoft"
	multiB[1][1] = "google"
	multiB[2][0] = "AT&T"
	multiB[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(multiB)

	as := [5]int{76, 77, 78, 79, 80}
	var bs []int = as[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(bs)

	cs := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(cs)

	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)]                                        //re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice), "\n")

	i := make([]int, 5, 5)
	fmt.Println(i)

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota", "Toyota1", "Toyota2", "Toyota3", "Toyota4")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside

	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

}

func countries() []string {
	countries := []string{"china", "japan", "usa"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}
