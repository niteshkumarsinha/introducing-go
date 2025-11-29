package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	defer fmt.Println("Program has ended. This wont be run.")
	fmt.Println("Hello, World!")
	fmt.Println("Enter your name.:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	println("Hello, " + input)

	os.Exit(0)
}
