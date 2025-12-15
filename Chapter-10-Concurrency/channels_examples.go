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

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
