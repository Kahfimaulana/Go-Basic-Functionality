package main

import "fmt"

func findGreatestNumber(args ...float64) float64{
	
	min := args[0]
	var result float64
	
	for _, v := range args {

	    if min < v {
	    	min = v
		result = min
	    }
	}

	return result
}

func main() {
    
    fmt.Println(findGreatestNumber(6,7,8,1,45,7,43))
}
