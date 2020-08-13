// goroutine dealing with race condition using channels (fatal map read write)
package main

import (
	"fmt"
	"time"
)

var MAPPER = map[string]int{
	"FOO": 0,
	"BAR": 0,
}

func handleMapper(channel chan map[string]int) {
	for {
		select {
		case data := <-channel:
			for k, v := range data {
				MAPPER[k] = v
			}
		case channel <- MAPPER:
		}
	}
}

func write(key string, value int, channel chan map[string]int) {
	channel <- map[string]int{key: value}
}

func read(key string, channel chan map[string]int) int {
	data := <-channel
	return data[key]
}

func main() {
	start := time.Now()
	channel := make(chan map[string]int)

	go handleMapper(channel)
	for i := 0; i < 1000; i++ {
		go write("FOO", i, channel)
		go write("BAR", i, channel)
		go read("FOO", channel)
		go read("BAR", channel)
	}

	elapsed := time.Since(start)
	fmt.Printf("FOO: %d\n", read("FOO", channel))
	fmt.Printf("BAR: %d\n", read("BAR", channel))
	fmt.Printf("Took %s\n", elapsed)
}
