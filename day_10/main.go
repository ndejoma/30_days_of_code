package main

import "fmt"

func main() {
	n := 3
	fmt.Println(fibonacci(n)) //6
}

func fibonacci(n int) int {
	if n == 1 {
		return n
	}
	return n * fibonacci(n-1)
}
