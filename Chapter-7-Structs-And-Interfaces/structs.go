package main


import (
	"fmt"
	"math"
)

const circleAreaFormat = "Circle Area: %.2f\n"

type Circle struct {
	x, y, r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	var c1 Circle

	c2 := new(Circle)
	c2.x = 1
	c2.y = 2
	c2.r = 3

	fmt.Println(c1)
	// Calculate circle area
	c3 := Circle{0, 0, 5}
	area := math.Pi * c3.r * c3.r
	fmt.Println(c3)
	fmt.Printf(circleAreaFormat, area)
	fmt.Println(c3)
	var c4 *Circle = &Circle{0, 0, 10}
	area2 := math.Pi * c4.r * c4.r
	fmt.Println(c4)
	fmt.Printf(circleAreaFormat, area2)
	fmt.Println(c4)
	fmt.Printf("Circle Area: %.2f\n", area2)

	fields_example()
	method_example()
}

func fields_example() {
	var c *Circle = new(Circle)
	c.x = 10
	c.y = 20
	fmt.Println(c)
	fmt.Printf(circleAreaFormat, circleArea(c))

	var c2 Circle = Circle{0, 0, 5}
	fmt.Println(c2)
	fmt.Printf(circleAreaFormat, circleArea(&c2))
	fmt.Println(c2)
	fmt.Println("Circle Area: %.2f\n", circleArea(&c2))
}

func method_example() {
	c := Circle{0, 0, 5}
	fmt.Println(c)
	fmt.Printf(circleAreaFormat, c.Area())

	c2 := &Circle{0, 0, 10}
	fmt.Println(c2)
	fmt.Printf(circleAreaFormat, c2.Area())
}