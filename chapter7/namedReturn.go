package main

import "fmt"

func f2() (r int){
   r = 100
   return
}

func main(){
   x := f2()
   fmt.Println(x)
}
