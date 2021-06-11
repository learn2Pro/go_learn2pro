package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}

//func createQuery(q interface{}) {
//	t := reflect.TypeOf(q)
//	v := reflect.ValueOf(q)
//	k := t.Kind()
//	fmt.Println("Kind ", k)
//	fmt.Println("Type ", t)
//	fmt.Println("Value ", v)
//
//}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)

}
