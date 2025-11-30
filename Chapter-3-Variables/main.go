package main

import "fmt"

var x string = "Global Variable"

var (
	a = 5
	b = 10
	c = 15
)

func main(){
	var x string = "Hello, World!"
	x = "First Go Program"
	fmt.Println(x)
	x = "Hello, Go!"
	fmt.Println(x)
	f()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	g()

	y := 5
	y += 1
	fmt.Println(y)
	fmt.Println(fahrenheitToCelsius(98.6))
	fmt.Println(celsiusToFahrenheit(37))

	fmt.Println(feetIntoMeters(5.5))
	fmt.Println(metersIntoFeet(1.6764))
}

func f(){
	fmt.Println(x)
}

func g(){
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println("Output:", output)
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func feetIntoMeters(feet float64) float64 {
	return float64(feet) * 0.3048
}

func metersIntoFeet(meters float64) float64 {
	return meters / 0.3048
}