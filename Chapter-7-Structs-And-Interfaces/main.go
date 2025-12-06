package main

import (
	"fmt"
	"math"
)


func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x2, y1)
	w := distance(x1, y1, x1, y2)
	return l * w
}

func circleArea(x, y, radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	dist := distance(1, 2, 4, 6)
	area := rectangleArea(1, 2, 4, 6)
	circle := circleArea(0, 0, 5)

	fmt.Printf("Distance: %.2f\n", dist)
	fmt.Printf("Rectangle Area: %.2f\n", area)
	fmt.Printf("Circle Area: %.2f\n", circle)

}

// This program finds the area of a rectangle and a circle. Keeping track of all the coor‐
// dinates makes it difficult to see what the program is doing and will likely lead to mis‐
// takes.

// In the next chapter, you'll learn how to use structs to group related data together,
// making your code easier to read and maintain.

