package main

import "fmt"

func main(){
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	val1, ok1 := x["key"]
	fmt.Println("The value:", val1, "Present?", ok1)
	delete(x, "key")
	fmt.Println(x["key"])

	val, ok := x["key"]
	fmt.Println("The value:", val, "Present?", ok)

	printElements()

	x1 := [6]string{"a","b","c","d","e","f"}
	fmt.Println(x1[2:5])

	x2 := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	min := x2[0]
	for _, v := range x2 {
		if v < min {
			min = v
		}
	}
	fmt.Println("The minimum value is:", min)


}


func printElements(){
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
	}

	for k, v := range elements {
		fmt.Println(k, "->", v["name"], "->", v["state"])
	}


}
