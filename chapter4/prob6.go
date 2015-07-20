package main 

import "fmt"

func convert(x float64){
	
	var result = x * 0.3048

	fmt.Println("The result is : ", result)
}

func main(){
	fmt.Println("Input feet : ")
	
	var input float64

	fmt.Scanf("%f", &input)

	convert(input)
}
