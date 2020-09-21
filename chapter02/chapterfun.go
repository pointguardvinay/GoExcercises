package main

import (
	"fmt"
)

var y int

func increment() int {
	y++
	return y
}

func main() {

	fmt.Println("Hello World ")
	var x int
	x = 0
	greeting := func() {

		x++
	}

	greeting()
	fmt.Println(x)
	greeting()
	fmt.Println(x)
	y = 0
	fmt.Println(y)
	increment()
	fmt.Println(y)
	increment()
	fmt.Println(y)

}
