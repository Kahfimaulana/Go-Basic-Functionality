package main

import "fmt"

func f2() (int, int){
    return 5, 10
}

func main(){
    x, y := f2()

    fmt.Println(x, y)
}
