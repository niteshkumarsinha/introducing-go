package main

import "fmt"


func main() {
	numbers := []float64{90, 80, 70, 60, 50}
	average(numbers)

	fmt.Println(f())
	fmt.Println(add(1, 2, 3, 4, 5))
	fmt.Println(add(10, 20))
	PrintLn(1, 2, 3, 4, 5)


	add := func(a, b int) int {
		return a + b
	}

	fmt.Println("The sum is:", add(3, 4))

	increment := incrementer()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	even := makeEvenGenerator()

	for i := 0; i < 100; i++ {	
		fmt.Println(even())
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
	}

}

func average(numbers []float64) {
	total := 0.0

	for _, number := range numbers {
		total += number
	}

	avg := total / float64(len(numbers))
	fmt.Printf("The average is: %.2f\n", avg)
}


func f() (int, int){
	fmt.Println("Hello from function f")
	return 5, 10
}

func f2(){
	fmt.Println("Hello from function f2")
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func PrintLn(a ...interface{}) (n int, err error) {
	xs := []int{1, 2, 3}
	return fmt.Println(add(xs...))
}

func incrementer() func() int {
	x := 0
	increment := func() int {
		x++
		return x
	}
	return increment
}

func makeEvenGenerator() func() int {
	i := 0
	return func() int {
		i += 2
		return i
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

