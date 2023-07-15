package main

import "fmt"

// CHALLENGE: day 3 HackerRank: Intro to Conditional Statements
/**
Task
Given an integer,n , perform the following conditional actions:

If n is odd, print Weird
If n is even and in the inclusive range 2 of 5 to , print "Not Weird"
If n is even and in the inclusive range of 6 to 20, print "Weird"
If n is even and greater than 20 , print "Not Weird"
*/
// func main() {
// 	var evenNo int32 = 4
// 	printMessage(evenNo)
// }

// func isEven(n int32) bool {
// 	return n%2 == 0
// }

// func printMessage(n int32) {
// 	if isOdd(n) {
// 		fmt.Println("Weird")
// 	}
// 	if isEven(n) {
// 		if n >= 2 && n <= 5 {
// 			fmt.Println("Not weird")
// 		}
// 		if n >= 6 && n <= 20 {
// 			fmt.Println("Weirdooooo")
// 		}
// 		if n > 20 {
// 			fmt.Println("Not Weird")
// 		}
// 	}
// }

//working solution

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	if n%2 != 0 {
		fmt.Println("Weird")
	} else {
		//by now it is an even number
		if n >= 2 && n <= 5 {
			fmt.Println("Not Weird")
		} else if n >= 6 && n <= 20 {
			fmt.Println("Weird")
		} else {
			fmt.Println("Not Weird")
		}
	}
}
