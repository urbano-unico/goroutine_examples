// goroutine using mutex to avoid race condition
package main

import (
	"fmt"
	"sync"
	"time"
)

var MUX = sync.Mutex{}
var MAPPER = map[string]int{
	"FOO": 0,
	"BAR": 0,
}

func write(key string, value int) {
	MUX.Lock()
	MAPPER[key] = value
	MUX.Unlock()
}

func read(key string) int {
	MUX.Lock()
	defer MUX.Unlock()

	return MAPPER[key]
}

func main() {
	start := time.Now()

	for i := 0; i < 1000; i++ {
		go write("FOO", i)
		go write("BAR", i)
		go read("FOO")
		go read("BAR")
	}

	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
}
