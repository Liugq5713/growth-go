package main

import (
	"fmt"
)

// Retriver test
type Retriver interface {
	get(string) string
}

func main() {
	test := test()
	fmt.Println("hi", test)
}
