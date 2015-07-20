package main

import "fmt" 

func smallestNumber(args []int) int{
	small := 0

	for i := 0;i < len(args); i++{
	  for j := 0; j < len(args); j++{
	    x := args[i];
	     x2 := args[j];
	     if x2 > x {
	      small = x2
	    } else {
	      small = x
	    }
	  }
		
	}

	return small
}

func searchSmallest(args []int) int{
	result := args[0]
	for _, v := range args{
		if v < result {
			result = v
		}
	}

	return result
}

func main(){
	 x := []int{
	 	48,96,86,68,9,
	 }
	fmt.Println(smallestNumber(x))
	fmt.Println(searchSmallest(x))
}
