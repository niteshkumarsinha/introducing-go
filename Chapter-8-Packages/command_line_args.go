package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// Define flags
	maxp := flag.Int("maxp", 10, "Maximum number of connections")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
