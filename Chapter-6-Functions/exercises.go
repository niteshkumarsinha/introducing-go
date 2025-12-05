package main


import "fmt"

func sum(a int, b int) int {
	return a + b
}

func half(n int) (int, string) {
	if n%2 == 0 {
		return n / 2, "Even"
	} else {
		return n / 2, "Odd"
	}
}

func greatest(a ...int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func oddNumberGenerator() func() int {
	n := -1
	return func() int {
		n += 2
		return n
	}
}

func fib(n int) int {
	if n < 0 {
		return -1
	}
	
	if n == 0 || n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func square(x *float64){
	*x = *x * *x
}

func swap(a, b *int){
	*a, *b = *b, *a
}

func main(){
	fmt.Println("Sum of 3 and 5 is:", sum(3, 5))
	num, isEvOrOdd := half(7)
	fmt.Println("7 is", isEvOrOdd)
	fmt.Println("Half of 7 is:", num)
	num, isEvOrOdd = half(10)
	fmt.Println("10 is", isEvOrOdd)
	fmt.Println("Half of 10 is:", num)

	fmt.Println("Greatest among 3, 5, 2, 8, 1 is:", greatest(3, 5, 2, 8, 1))

	oddGen := oddNumberGenerator()
	for i := 0; i < 5; i++ {
		fmt.Println("Next odd number is:", oddGen())
	}

	fmt.Println("Fibonacci of 6 is:", fib(6))
	fmt.Println("Fibonacci of 10 is:", fib(10))

	defer func(){
		e := recover()
		if e != nil {
			fmt.Println("Recovered in main:", e)
		}
	}()

	fmt.Println(new(int))

	x := new(int)
	*x = 42
	fmt.Println(*x)	

	y := 3.0
	square(&y)
	fmt.Println("Square is:", y)

	fmt.Println("End of main function")

	a := 1
	b := 2
	swap(&a, &b)
	fmt.Println("After swapping, a =", a, "b =", b)

	panic("Panic in main function")

	

}