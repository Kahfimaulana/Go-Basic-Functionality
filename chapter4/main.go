package main

import "fmt"

func main(){
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
	x = x + "third"
	fmt.Println(x)
	var(
		a=10
		b=20
		c=30
	)
	fmt.Println("a + b + c =", a+b+c)
}
