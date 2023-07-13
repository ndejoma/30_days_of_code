package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var myInt int
	var myDouble float64
	var myString string

	// Read and save an integer, double, and String to your variables.
	count := 0
	for scanner.Scan() {
		switch count {
		case 0:
			myInt, _ = strconv.Atoi(scanner.Text())
		case 1:
			myDouble, _ = strconv.ParseFloat(scanner.Text(), 64)
		case 2:
			myString = scanner.Text()
		}
		count++
	}
	// Print the sum of both integer variables on a new line.
	fmt.Println(myInt + int(i))
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", myDouble + d)
	// Concatenate and print the String variables on a new line
	fmt.Println( s + myString)
	// The 's' variable above should be printed first.

}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	count := 0
// 	var str1 string
// 	var str2 string
// 	var str3 string
// 	for scanner.Scan() {
// 		if count > 2 {
// 			break
// 		}
// 		switch count {
// 		case 0:
// 			str1 = scanner.Text()
// 		case 1:
// 			str2 = scanner.Text()
// 		case 2:
// 			str3 = scanner.Text()
// 		}
// 		count++
// 	}
// 	fmt.Println(str1)
// 	fmt.Println(str2)
// 	fmt.Println(str3)
// }
