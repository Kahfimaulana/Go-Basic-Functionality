package main 

import "fmt"

func square(x *float64){
	*x = *x * *x
}

func main(){
	var x float64
	x = 15
	square(&x)
	fmt.Println(x)
}
