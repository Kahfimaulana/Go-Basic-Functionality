package main

import "fmt"

func swap(x *int, y *int){
	*x, *y = *y, *x
}

func main(){
	x := 50
	y := 10
	swap(&x, &y)
	fmt.Println(x, y)
}
