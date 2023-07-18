package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Create a new reader, reading from Stdin
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	if err != nil {
		log.Fatal(err)
	}
	var evenString string
	var oddString string
	if len(input) > 1 {
		for i := 0; i < len(input); i++ {
			if i == 0 || i%2 == 0 {
				evenString += string(input[i])
			} else {
				oddString += string(input[i])
			}
		}
		fmt.Println(evenString + " " + oddString)
	}
}


//solution hackerRank
// package main

// import (
//     "fmt"
//     "os"
// )

// func main() {
//     var testCases int
//     fmt.Scanf("%d", &testCases) // This reads the number of test cases

//     for i := 0; i < testCases; i++ {
//         var input string
//         fmt.Scanf("%s", &input) // This reads each test case

//         var evenString, oddString string
//         for j := 0; j < len(input); j++ {
//             if j%2 == 0 {
//                 evenString += string(input[j])
//             } else {
//                 oddString += string(input[j])
//             }
//         }

//         fmt.Println(evenString + " " + oddString) // This prints the result
//     }
// }
