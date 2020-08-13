// goroutine using WaitGroup to avoid race condition (sync)
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

func syncWrite(key string, value int, wg *sync.WaitGroup) {
	wg.Add(1)
	go write(key, value, wg)
	wg.Wait()
}

func syncRead(key string, wg *sync.WaitGroup) {
	wg.Add(1)
	go read(key, wg)
	wg.Wait()
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
		syncWrite("FOO", i, &wg)
		syncWrite("BAR", i, &wg)
		syncRead("FOO", &wg)
		syncRead("BAR", &wg)
	}

	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}
