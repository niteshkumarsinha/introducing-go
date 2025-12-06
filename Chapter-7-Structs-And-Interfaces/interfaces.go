package main


import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

type Rectangle struct {
	l float64
	w float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	return r.l * r.w
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	return 2 * (r.l + r.w)
}


// func totalArea(circles ...Circle) float64 {
// 	var total float64
// 	for _, c := range circles {
// 		total += c.area()
// 	}
// 	return total
// }

// func totalArea(circles []Circle, rectangles []Rectangle) float64{
// 	var total float64
// 	for _, c := range circles {
// 		total += c.area()
// 	}
// 	for _, r := range rectangles {
// 		total += r.area()
// 	}
// 	return total
// }


func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64{
	totalArea := 0.0
	for _, s := range m.shapes {
		totalArea += s.area()
	}
	return totalArea
}

func (m *MultiShape) perimeter() float64 {
	var totalPerimeter float64
	for _, s := range m.shapes {
		totalPerimeter += s.perimeter()
	}
	return totalPerimeter
}


func main(){
	area := totalArea(&Rectangle{10, 10}, &Circle{10})
	fmt.Println(area)

	multiShape := MultiShape{
		shapes : []Shape {
			&Circle{5},
			&Rectangle{10, 10},
		},
	}

	fmt.Println("The area of multishapes:", multiShape.area())
	fmt.Println("Permimeter of multishapes:", multiShape.perimeter())
}