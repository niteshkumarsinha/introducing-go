package main

import (
	"os"
	"fmt"
)

func first() {
	fmt.Println("First function called")
}

func second() {
	fmt.Println("Second function called")
}

func main() {
	defer second()
	first()

	f, _ := os.Open("non_existent_file.txt")
	defer f.Close()

	fmt.Println("End of main function")

	mainPanicRecover()


	// -- 

	defer func(){
		str := recover()
		if str != nil {
			fmt.Println("Recovered in anonymous function:", str)
		}
	}()

	panic("Panic in main function")
	
}

func mayPanic() {
	panic("A panic occurred!")	
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func mainPanicRecover() {
	defer recoverFromPanic()
	mayPanic()
	fmt.Println("This line will not be executed if panic occurs")
}
