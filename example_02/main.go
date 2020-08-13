// goroutine example waiting goroutine to be executed
package main

import (
	"fmt"
	"time"
)

func print() {
	fmt.Println("Printing from goroutine")
}

func main() {
	go print()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Printing from main")
}
