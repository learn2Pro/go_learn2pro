package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}
type MyString string

func (str MyString) FindVowels() []rune {
	var vowels []rune
	for _, c := range str {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels = append(vowels, c)
		}
	}
	return vowels
}

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}


func main() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())

	var s interface{} = 56
	assert(s)
	var i interface{} = "Steven Paul"
	assert(i)
}
