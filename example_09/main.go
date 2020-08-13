// goroutine race condition error
package main

import (
	"fmt"
	"time"
)

var MAPPER = map[string]int{
	"FOO": 0,
	"BAR": 0,
}

func write(key string, value int) {
	MAPPER[key] = value
}

func read(key string) int {
	return MAPPER[key]
}

func main() {
	start := time.Now()

	for i := 0; i < 50; i++ {
		go write("FOO", i)
		go write("BAR", i)
		go read("FOO")
		go read("BAR")
	}

	elapsed := time.Since(start)
	time.Sleep(1 * time.Second)
	fmt.Printf("Took %s\n", elapsed)
}
