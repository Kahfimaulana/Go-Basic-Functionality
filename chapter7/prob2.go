package main

import "fmt"

func main() {
    result := 0
    half := func(x int) (int, bool) {
    	 result = x / 2

         if result % 2 ==  1 {
             return result, false

	 }else if result % 2 == 0 {
	     return result, true
	 }else{
	     return 1, true
	 }
    }

    fmt.Println(half(7))
}
