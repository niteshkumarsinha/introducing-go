package main

import (
	"fmt"
	"packages_example/maths"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	fmt.Println(xs)
	avg := maths.Average(xs)
	fmt.Println(avg)
}
