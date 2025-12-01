package main

import "fmt"

func main() {
	//arrayExample()
	arrayExample1()
	//arrayExample2()
}	

func arrayExample1() {
	var total float64 = 0
	for i, value := range [5]float64{98, 93, 77, 82, 83} {
		total += value
		fmt.Println("Index:", i, "Value:", value)
	}
	fmt.Println("Total:", total)
	fmt.Println("Average:", total/5)	
}

func arrayExample2() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println("Total:", total)
	fmt.Println("Average:", total/float64(len(x)))

	scores := [5]float64{98, 93, 77, 82, 83}
	average := computeAverage(scores[:])
	fmt.Println("Average from function:", average)

	scores2 := []float64{98, 93, 77, 82, 83, 100, 67, 89}
	average2 := computeAverage(scores2)
	fmt.Println("Average from function with slice:", average2)	
}

func computeAverage(scores []float64) float64 {
	var total float64 = 0
	for i := 0; i < len(scores); i++ {
		total += scores[i]
	}
	return total / float64(len(scores))
}

func arrayExample() {
	// Declare an array of integers with a fixed size of 5
	var x [5]int
	x[4] = 100
	fmt.Println("Array x:", x)
	fmt.Println("Length of array x:", len(x))

	// Declare and initialize an array of integers
	numbers := [5]int{10, 20, 30, 40, 50}

	// Print the entire array
	fmt.Println("Array:", numbers)

	// Access and print individual elements
	fmt.Println("First element:", numbers[0])
	fmt.Println("Third element:", numbers[2])

	// Modify an element in the array
	numbers[1] = 25
	fmt.Println("Modified Array:", numbers)

	// Get the length of the array
	length := len(numbers)
	fmt.Println("Length of the array:", length)
}