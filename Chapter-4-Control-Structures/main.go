package main

import "fmt"

func main() {
	fizzBuzz()
	//divisibleByN(3, 100)
	// example()
	//printNums()
	//main2()
}

func main2(){
	fmt.Println(1)	
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("Infinite Loop")
		break
	}

	for k := 1; k <= 10; k++ {
		if k%2 == 0 {
			fmt.Println(k, "is even")
		} else {
			fmt.Println(k, "is odd")
		}
	}

	for n := 1; n <= 20; n++ {
		if n%3 == 0 && n%5 == 0 {
			fmt.Println(n, "FizzBuzz")
		} else if n%3 == 0 {
			fmt.Println(n, "Fizz")
		} else if n%5 == 0 {
			fmt.Println(n, "Buzz")
		} else {
			fmt.Println(n)
		}
	}
}

func printNums(){
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	for i := 0 ; i < 10; i++ {
		fmt.Println(i)
	}
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} 
	return false
}

func isOdd(n int) bool {
	if n%2 != 0 {
		return true
	} 
	return false
}

func printEvenOdd(){
	for i := 1; i <= 10; i++ {
		if isEven(i) {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}

func example(){
	var i int = 1
	for i <= 10 {
		if i % 2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
		i += 1
	}

	switch i {
	case 0: fmt.Println("i is zero")
	case 1: fmt.Println("i is one")
	case 2: fmt.Println("i is two")
	default: fmt.Println("i is greater than two")
	}

	i = 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}

func divisibleByN(n int, upToRange int) {
	i := 0
	for i <= upToRange {
		if i % n == 0 {
			fmt.Printf("%v is divisible by %v\n", i, n)
		}
		i++
	}
}

// fizzBuzz in range 1 to 100
func fizzBuzz(){
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}
}