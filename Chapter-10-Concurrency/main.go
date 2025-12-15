package main

import "fmt"
import "time"
import "math/rand"

func f(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
func main() {
	for i := 1; i <= 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
