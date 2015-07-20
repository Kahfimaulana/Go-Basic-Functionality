package main

import "fmt"

func main(){
    x := 0
    y := 1

    for i := 0; i < 10; i++{
	z := x + y
	fmt.Println(z)
	x = y
	y = z
    }
}
