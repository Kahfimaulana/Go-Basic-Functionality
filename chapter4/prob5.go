package main

import "fmt"

func convert(x float64){

	var C = (x - 32) * 5/9;

	fmt.Println("Celsius: ", C);
}

func main(){
	fmt.Println("Input the fahrenheit number : ")

	var input float64

	fmt.Scanf("%f", &input)

	fmt.Println("Fahrenheit : ", input)

	convert(input)
}
