// goroutine example `maybe it will execute`
package main

import (
	"fmt"
)

func print() {
	fmt.Println("Printing from goroutine")
}

func main() {
	go print()
	fmt.Println("Printing from main")
}
