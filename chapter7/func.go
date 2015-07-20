package main

import "fmt"

func f1() int{
    return f2()
}

func f2() int{
    return 1
}

func main(){
   
   fmt.Println(f1())

} 
