package main

import "fmt"

func main(){
	//add := func(x, y int) int{
	//     z := x + y
	//     return z
	//}

	//fmt.Println(add(1,2))

	nextEvent := retFunc()

	fmt.Println(nextEvent)
}

func retFunc() func() uint{
	u := uint(0)

	return func() (r uint){
		r = u
		u += 2
		return 
	}
}
