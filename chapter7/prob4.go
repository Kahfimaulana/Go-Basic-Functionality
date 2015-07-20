package main

import "fmt"

func makeOddGenerator() func() float64 {
	i := float64(1)
	return func() (r float64){
		r = i
		i += 2
		return
	}
}

func main(){
	fmt.Println(makeOddGenerator())
	fmt.Println(makeOddGenerator())
}
