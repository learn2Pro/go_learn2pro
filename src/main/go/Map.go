package main

import "fmt"

type employee struct {
	salary  int
	country string
}

func main() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	employeeSalary["steve"] = 12000
	employeeSalary["jamie"] = 15000
	employeeSalary["mike"] = 9000
	newEmp := "joe"
	v, ok := employeeSalary[newEmp]
	if ok {
		fmt.Println("Salary of", newEmp, "is", v)
		return
	}
	fmt.Println(newEmp, "not found")
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}

	fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)
	fmt.Println(employeeSalary)

	emp1 := employee{
		12,
		"abc",
	}
	emp2 := employee{
		salary:  14000,
		country: "Canada",
	}
	emp3 := employee{
		salary:  13000,
		country: "India",
	}
	employeeInfo := map[string]employee{
		"Steve": emp1,
		"Jamie": emp2,
		"Mike":  emp3,
	}

	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary:$%d  Country: %s\n", name, info.salary, info.country)
	}

	name := "Hello World"
	for _, c := range name {
		fmt.Printf("%x,%c\n", c, c)
	}

	h := "hello"
	fmt.Println(mutate([]rune(h)))

	a := 10
	var b *int = &a
	fmt.Printf("Type of a is %T\n", b)
	fmt.Println("address of b is", b)
	var emp4 Employee //zero valued struct
	fmt.Println("First Name:", emp4.firstName)
	fmt.Println("Last Name:", emp4.lastName)
	fmt.Println("Age:", emp4.age)
	fmt.Println("Salary:", emp4.salary)
	emp4.display()

}

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func (e Employee) display() {
	fmt.Printf("Salary of %s is %s%d", e.firstName, e.lastName, e.salary)
}

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}
