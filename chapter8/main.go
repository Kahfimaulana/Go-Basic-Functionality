package main

import "fmt"

func change(x *int){
	*x = 0
}

func main(){
	x := 5
	change(&x)
	fmt.Println(x)

	y := new(int)
	zero(y)
	fmt.Println(*y)
}

func zero(y *int){
	*y = 10
}
