package main

import "fmt"

func name(a string) {
    fmt.Println("hello, My name is", a + "\n")
    fmt.Println(len("hello, My name is"))
    fmt.Println("hello Kahfi"[1])
    fmt.Println("321325 * 424521 =", 321325 * 424521)
    fmt.Println((true && false) || (false && true) || !(false && false))
}

func main(){
	name("Kahfi Maulana Iskandar")
}
