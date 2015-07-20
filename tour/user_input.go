package main

import "fmt"

// func main() {
	// fmt.Println("input the number")
	// var input float64
	// fmt.Scanf("%f", &input)
	// output := input * 2
	// fmt.Println("the result is ", output)
// }

// func main() {
// 	fmt.Println("This program to convert Fahrenheit to Celcius")
// 	fmt.Println("Enter the Fahrenheit value: ")
// 	var input float64

// 	fmt.Scanf("%f", &input)
// 	x := (input - 32) * 5/9
// 	fmt.Println(x)
// }

var feet float64 = 0.3048

func main() {
	var input float64
	fmt.Println("input feet :")
	fmt.Scanf("%f", &input)
	output := input * feet
	fmt.Println(output)
}