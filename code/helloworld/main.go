package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	foo(50)
	bar(20)
}

func foo(n int) {
	if n > 10 {
		fmt.Println("greater")
		return
	}
	fmt.Println("less")
}

func bar(n int) {
	val := 3
	if n > 10 {
		val = 10 * 12
	}
	fmt.Println(val)
}
