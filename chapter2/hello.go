package main

import "fmt"

func name(s string){
	fmt.Println("hello my name is ", s)
	fmt.Println("Hello Word"[1])
	fmt.Println(len("hello my name is"))
	fmt.Println("hello" + "kahfi")
}

func main(){
	name("Kahfi Maulana")
}
