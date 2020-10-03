package main

import (
	"fmt"
)

func main() {

	str1 := "This"
	str2 := "That"

	r1 := []rune(str1)
	r2 := []rune(str2)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r1)

	if len(r1) != len(r2) {
		fmt.Println("Length is similar")
	} else {

	}

	for _, v := range r1 {

		fmt.Println(v, "And value is ", v)
	}

	// map declaration

	m := map[string]int{"First": 1, "Second": 2, "Third": 3}

	//accessing the elements
	fmt.Println("1===>", m["First"])
	//updating the value
	m["First"] = 10
	//deleting the element
	delete(m, "Third")
	for k, v := range m {

		fmt.Println(k, " => ", v)
	}

	// Slice & Array
	// array declaration with fixed array size
	var a [4]int
	fmt.Println(a)
	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println(c)
	// slice declaration without size... size is not necessary in Slice.
	var b []int
	fmt.Println(b)

}
