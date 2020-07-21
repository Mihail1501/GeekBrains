package main

import (
	"fmt"
)

func fibi(n int) int {
	var a, b int = 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

type fibfunc func(int) int

func printFib(fib fibfunc, a, b int) {
	for i := a; i < b; i++ {
		fmt.Printf("%d, ", fib(i))
	}
	fmt.Println("...")
}

func main() {
	printFib(fibi, 0, 16)
}
