package main

import (
	"bufio"
	"fmt"
	"os"
)

func doSomeChoice(){
	reader := bufio.NewReader(os.Stdin)
	char,_,err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(char)
	switch char {
	case 'A':
		fmt.Println("A key")
		break
	case 'B':
		fmt.Println("B key")
		break
	}
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	scanner()
	doSomeChoice()
	fmt.Println("hello world")
}