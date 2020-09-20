package main

import (
	"fmt"
	"strconv"

	"github.com/goexcercises/chapter01/stringutil"
)

const (
	c0 = iota
	c1 = 1 + iota
	c2 = 1 + iota
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

	//checking the redeclaration you need a new variable with it.

	n, a := 10, 30
	//a := 20
	fmt.Println(a, n)

	//defining your own type

	type myType int

	var ab myType
	ab = 10
	fmt.Println("This is my type", ab)
	fmt.Printf("%T", ab)
	a = int(ab)
	fmt.Println(ab)
	fmt.Printf("%T", a)
	fmt.Printf("%T\n\n", ab)

	//----------------------------------------------------------
	fmt.Println(c0, c1, c2)

	var testV float64
	testV = 1234567890.123456
	fmt.Printf("format printing ===%20.2f\n\n", testV)

	// string slicing

	var abc string = "This is my string"

	fmt.Println(abc[1:5])
	fmt.Println(abc[:5])

	var testStr string = "56.33"
	answer, _ := strconv.Atoi(testStr)
	fmt.Println(answer)

	fmt.Println("Hello, playground")

	//slice definition shorthand

	myintSlice := []int{1, 2, 3, 4, 5}
	mystringSlice := []string{"First", "Second", "Third", "Fourth", "Fifth", "Sixth", "Seventh"}

	// Making a slice from the make keyword

	mymakeSlice := make([]int, 3, 6)
	mymakeSlice[0] = 1
	mymakeSlice[1] = 2
	mymakeSlice[2] = 3
	//mymakeSlice[3] = 4  //Throws an error

	//

	fmt.Println("Printing MyintSlice having capacity", cap(myintSlice), "\tand Length", len(myintSlice), "\tPrinting it", myintSlice)
	fmt.Println("Printing mystringSlice having capacity", cap(mystringSlice), "\tand Length", len(mystringSlice), "\tPrinting it", mystringSlice)
	fmt.Println("Printing mymakeSlice having capacity", cap(mymakeSlice), "\tand Length", len(mymakeSlice), "\tPrinting it", mymakeSlice)

	//slicing the slice

	mysSlice := myintSlice[:3] //inherits the capacity from the myintSlice
	fmt.Println("Printing mysSlice having capacity", cap(mysSlice), "\tand Length", len(mysSlice), "\tPrinting it", mysSlice)

	//appending the slice //we can append beyond its capacity
	mymakeSlice = append(mymakeSlice, 4) // APPENDING BEYOND ITS CAPACITY
	mymakeSlice = append(mymakeSlice, 5, 6, 7, 8)
	fmt.Println("Printing mymakeSlice having capacity", cap(mymakeSlice), "\tand Length", len(mymakeSlice), "\tPrinting it", mymakeSlice)

	//using append you can delete the elements

	mymakeSlice = append(mymakeSlice[:2], mymakeSlice[4:]...)
	fmt.Println("Printing newly appended mymakeSlice having capacity", cap(mymakeSlice), "\tand Length", len(mymakeSlice), "\tPrinting it", mymakeSlice)

	// M A P
	myMap := map[string]int{"Vinay": 38, "Pruthvi": 11}
	fmt.Println("Priting myMap = ", myMap)

	myMap["newElement"] = 43 //adding new element

	fmt.Println("Priting myMap = ", myMap)

	myMap["newElement"] = 54 //updating an entry

	fmt.Println("Priting myMap = ", myMap)

	delete(myMap, "newElement") //deleting an entry
	fmt.Println("Priting myMap = ", myMap)

	for key, value := range myMap {
		fmt.Println("Printing from the for range", key, value)
	}
	myMap1 := new(map[string]int)
	fmt.Println("Created the map by using new keyword = ", myMap1)

	fmt.Println(myMap1)
	fmt.Println(*myMap1)
	*myMap1 = make(map[string]int)
	(*myMap1)["This is first"] = 20
	//&myMap1["This is Second"] =3
	//myMap1["This is Third"] =3
	//myMap1["This is Fourth"] =2

	fmt.Println("Created the map by using new keyword = ", *myMap1)

	var mySlice7 *[]int = new([]int)
	(*mySlice7)[0] = 1

	fmt.Println(*mySlice7)

}
