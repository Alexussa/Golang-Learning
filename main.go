package main

import (
	"fmt"
)

func main() {
	// Same package principal
	//	samePackage.SamePackage()

	// Multiple returns
	//	fmt.Println(returnMultiple("Jane ", "Doe "))

	// variadic params
	n := variadicParams(43, 56, 87, 12, 45, 57)
	fmt.Println(n)

	// variadic args
	data := []float64{43, 56, 87, 12, 45, 57}
	x := variadicParams(data...)
	fmt.Println(x)

	// The difference between the two is how you pass the params
	// in the first example you can pass as many params as you want
	// in the second example you need to pass a valid slice type eg.[]float64{1, 2, 3}
}

func utf8() {
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}

func varZeroValue() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println()
}

// return multiple
func returnMultiple(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

// variadic params
func variadicParams(sf ...float64) float64 {
	// this function gets the average of all numbers passed
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}