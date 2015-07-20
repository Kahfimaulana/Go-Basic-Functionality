package main 

import "fmt"

func main(){
	var arr [5]float64

	arr[0] = 98
	arr[1] = 93
	arr[2] = 77
	arr[3] = 82
	arr[4] = 83
	
	var total float64 = 0

	for i:= 0 ; i<= 4 ; i++{
		total += arr[i]
	}

	fmt.Println(total/5)
}
