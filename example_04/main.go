// goroutine waiting execution with WaitGroup (sync)
package main

import (
	"fmt"
	"sync"
	"time"
)

const PRINTS = 5

func print(i int, wg *sync.WaitGroup) {
	fmt.Printf("Printing from goroutine %d\n", i)
	wg.Done()
}

func main() {
	start := time.Now()

	wg := sync.WaitGroup{}
	for i := 0; i < PRINTS; i++ {
		wg.Add(1)
		go print(i, &wg)
		wg.Wait()
	}
	fmt.Println("Printing from main")

	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}
