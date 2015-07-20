package main

import "fmt"

func main(){
	
	/*create slice*/
	//arr := []float64{1,2,3,4,5,6}
	//x := arr[0:]

	//fmt.Println(arr)
	//fmt.Println(x)

	/*append slice*/
	//arr := []float64{1,2,3}
	//arr2 := append(arr, 4,5)

	//fmt.Println(arr2)

	/*Copy slice*/
	arr := []int{1,2,3,4,5,6}
	arr2 := make([]int, 3)
	copy(arr2, arr)

	fmt.Println(arr, arr2)
}
