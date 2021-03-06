// goroutine waiting execution with channel (end unexpected)
package main

import (
	"fmt"
	"time"
)

const PRINTS = 5

func print(i int, channel chan bool) {
	fmt.Printf("Printing from goroutine %d\n", i)
	channel <- true
}

func readChannel(channel chan bool) {
	for {
		select {
		case <-channel:
		default:
			return
		}
	}
}

func main() {
	start := time.Now()
	channel := make(chan bool)

	for i := 0; i < PRINTS; i++ {
		go print(i, channel)
	}
	readChannel(channel)

	fmt.Println("Printing from main")

	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}
