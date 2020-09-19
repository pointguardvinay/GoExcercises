package main

import (
	"fmt"

	"github.com/goexcercises/chapter01/stringutil"
)

func main() {

	var message string
	message = "This is my vairble of type string"
	var a, b int
	var c bool
	var d, e float64

	fmt.Println("==>", message)
	fmt.Println(a, b, c)
	fmt.Println(d, e)

	fmt.Println("THis is my first program")
	v := stringutil.Reverse("It May Fail")
	fmt.Println("Revse is ", v)

	//checking the redeclaration

	n, a := 10, 30
	//a := 20
	fmt.Println(a, n)
}
