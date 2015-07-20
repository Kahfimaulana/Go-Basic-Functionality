package main

import "fmt"

func average(arr[] float64){
	
	var total float64

	for i := 0; i < len(arr) ; i++{
		total += arr[i]
	}

	fmt.Println(total / float64(len(arr)))
}

func average2(arr[] float64){
	var total float64

	for _, value := range arr{
		total += value 
	}

	fmt.Println(total / float64(len(arr)))
}

func main(){
	x := []float64{
	30,
	40,
	56,
	90,
	70,
	}

	average(x)

	average2(x)
}
