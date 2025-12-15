package main

import (
	"fmt"
	"time"
)

//Channels provide a way for two goroutines to communicate with each other
//and synchronize their execution. Here is an example program using channels:

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func select_example() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from c1"
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			c2 <- "from c2"
			time.Sleep(3 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			default:
				fmt.Println("default")
			}
		}
	}()
}

func main() {
	//var c chan string = make(chan string)
	//go pinger(c)
	//go ponger(c)
	//go printer(c)
	select_example()
	var input string
	fmt.Scanln(&input)
}
