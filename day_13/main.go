package main

import "fmt"

func main() {
	name := "obina"

	for char, v := range name {
		fmt.Println(char, v)
	}
}
