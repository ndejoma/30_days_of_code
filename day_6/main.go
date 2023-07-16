package main

import "fmt"

func main() {
	n := 3
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
}
