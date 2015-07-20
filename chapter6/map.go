package main

import "fmt"

func main(){
	
	x := make(map[string]string)
	x["H"] = "Hydrogen"
	x["He"] = "Helium"
	x["Li"] = "Litium"
	x["Be"] = "Beryllium"
	x["B"] = "Boron"
	x["C"] = "Carbon"
	x["N"] = "Nitroge"
	x["O"] = "Oxygen"
	x["F"] = "Fluorine"
	x["Ne"] = "Neon"

	fmt.Println(x["Li"])

	name, ok := x["Un"]
	fmt.Println(name, ok)
	
	//if name, ok := x["Un"]; ok{
	//	fmt.Println(name,ok)
	//}
}
