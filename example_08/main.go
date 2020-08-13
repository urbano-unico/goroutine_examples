// goroutine waiting execution with channel
package main

import (
	"fmt"
	"time"
)

const PRINTS = 5

func print(channel chan bool) {
	for i := 0; i < PRINTS; i++ {
		go fmt.Printf("Printing from goroutine %d\n", i)
		channel <- true
	}
	close(channel)
}

func main() {
	start := time.Now()
	channel := make(chan bool)

	go print(channel)
	for range channel {
	}

	fmt.Println("Printing from main")

	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}
