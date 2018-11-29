package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Goroutine Tutorial", time.Second)

	go  compute(4)
	go compute(4)
	var input string
	// Anonymous Goroutine Functions
	go func() {
		fmt.Println("Executing my Concurrent anonymouse function")
	}()


	fmt.Scanln(&input)
}
