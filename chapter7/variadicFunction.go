package main

import "fmt"

func variadic(args ...int) int{
    result := 0

    for _, v := range args{
        result += v
    }

    return result
}

func main(){
    fmt.Println(variadic(1,2,3,4,5,5))
}
