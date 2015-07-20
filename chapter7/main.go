package main

import "fmt"

func average(xs []float64) float64{
    
    var result float64

    for _, value := range xs{
    	result += value
    }

    return result / float64(len(xs))
}

func main(){
    arr := []float64{90,87,68,79,90}

    fmt.Println(average(arr))
}
