// goroutine dealing with race condition using channels
package main

import (
	"fmt"
	"time"
)

var MAPPER = map[string]int{
	"FOO": 0,
	"BAR": 0,
}

func handleMapper(writter chan map[string]int, reader chan string, readerResult chan int) {
	for {
		select {
		case data := <-writter:
			for k, v := range data {
				MAPPER[k] = v
			}
		case key := <-reader:
			readerResult <- MAPPER[key]
		}
	}
}

func write(key string, value int, writter chan map[string]int) {
	writter <- map[string]int{key: value}
}

func read(key string, reader chan string, readerResult chan int) int {
	reader <- key
	return <-readerResult
}

func main() {
	start := time.Now()
	writter := make(chan map[string]int)
	reader := make(chan string)
	readerResult := make(chan int)

	go handleMapper(writter, reader, readerResult)
	for i := 0; i < 1000; i++ {
		go write("FOO", i, writter)
		go write("BAR", i, writter)
		go read("FOO", reader, readerResult)
		go read("BAR", reader, readerResult)
	}

	elapsed := time.Since(start)
	fmt.Printf("FOO: %d\n", read("FOO", reader, readerResult))
	fmt.Printf("BAR: %d\n", read("BAR", reader, readerResult))
	fmt.Printf("Took %s\n", elapsed)
}
