// package main

// import (
//     "fmt"
//     "math"
// )

// func isPrime(n int32) bool {
//     if n <= 1 {
//         return false
//     }
//     sqRoot := int(math.Sqrt(float64(n)))
//     for i := 2; i <= sqRoot; i++ {
//         if n%int32(i) == 0 {
//             return false
//         }
//     }
//     return true
// }

//	func main() {
//	    fmt.Println(isPrime(2))  // true
//	    fmt.Println(isPrime(4))  // false
//	    fmt.Println(isPrime(5))  // false
//	    fmt.Println(isPrime(17)) // true
//	}
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	tests, err := strconv.ParseInt(input.Text(), 10, 64)
	if err != nil {
		panic("please enter a number")
	}
	for i := 0; i < int(tests); i++ {
		input.Scan()
		no, err := strconv.ParseInt(input.Text(), 10, 64)
		if err != nil {
			continue
		}
		if isPrime(int(no)) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	sqRoot := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqRoot; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
