package main

import "fmt"

func factorial(x float64) float64 {

	if x == 0 {
		return 1
	}

	return x * factorial(x-1)

}

func main() {

	fmt.Println("Input the number : ")

	var input float64

	fmt.Scanf("%f", &input)

	fmt.Println("The factorial from input number is : ", factorial(input))
}
