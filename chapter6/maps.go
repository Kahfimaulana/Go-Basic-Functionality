package main

import "fmt"

func main(){
	x := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		
		"HE": map[string]string{
			"name":"Helium",
			"state":"gas",
		},

		"Li": map[string]string{
			"name":"Lithium",
			"state":"solid",
		},

		"Be": map[string]string{
			"name":"Beryllium",
			"state": "solid",
		},

		"B": map[string]string{
			"name":"Boron",
			"state":"solid",
		},

		"C": map[string]string{
			"name":"Carbon",
			"state":"solid",
		},

		"N": map[string]string{
			"name":"Nitrogen",
			"state":"gas",
		},

		"O": map[string]string{
			"name":"Oxygen",
			"state":"gas",
		},

		"F": map[string]string{
			"name":"Fluorine",
			"state":"gas",
		},

		"Ne": map[string]string{
			"name":"Neon",
			"state":"gas",
		},
	}

	if el, ok := x["Li"]; ok{
		fmt.Println(el["name"], el["state"])
	}
}
