// goroutine using WaitGroup to avoid race condition
package main

import (
	"fmt"
	"sync"
	"time"
)

var MAPPER = map[string]int{
	"FOO": 0,
	"BAR": 0,
}

func write(key string, value int, wg *sync.WaitGroup) {
	MAPPER[key] = value
	wg.Done()
}

func read(key string, wg *sync.WaitGroup) int {
	defer wg.Done()

	return MAPPER[key]
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(4)
		go write("FOO", i, &wg)
		go write("BAR", i, &wg)
		go read("FOO", &wg)
		go read("BAR", &wg)
		wg.Wait()
	}

	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}
