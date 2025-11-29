package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =",1+1)
	fmt.Println(len("Hello, World!"))
	fmt.Println("Hello, World!"[1])
	fmt.Println("Hello, " + "World!")

	fmt.Println(true && false)
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(!false)
	fmt.Println(!true)

	fmt.Println(largestNDigitBinary(3))
	fmt.Println(largestNDigitBinary(4))
	fmt.Println(largestNDigitBinary(5))
	fmt.Println(largestNDigitBinary(8))

	fmt.Println(computeMultiplication(32132, 42452))

	fmt.Println(len("Nitesh Kumar"))
}

func computeMultiplication(a int, b int) int {
	return a * b
}

func largestNDigitBinary(n int) int {
	return pow(2, n) - 1
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}