// package main

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"strconv"
// )

// func main() {
// 	no := 125
// 	// for {
// 	// 	if no == 0 {
// 	// 		break
// 	// 	}
// 	// }
// 	// fmt.Println(no % 2)
// 	// fmt.Println(no / 5)
// 	binaryStr, err := decimalToBinary(no)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//count the number of maximum consecutive ones
// 	printMaxConsecutiveOnes(binaryStr)

// }

// func printMaxConsecutiveOnes(binaryStr string) {
// 	max := 0
// 	currMax := 0
// 	for i := 0; i < len(binaryStr); i++ {
// 		if string(binaryStr[i]) == "0" {
// 			//we have finished counting  the 1'S
// 			currMax = 0
// 		} else {
// 			currMax++
// 			if currMax > max {
// 				max = currMax
// 			}
// 		}
// 	}
// 	fmt.Println(max)
// }

// func decimalToBinary(no int) (string, error) {
// 	if no < 0 {
// 		return "", errors.New("input number must be a positive integer")
// 	}
// 	const divisor = 2
// 	dividend := no
// 	var binaryStr string
// 	for dividend > 0 {
// 		quotient := dividend / divisor
// 		remainder := dividend % divisor
// 		binaryStr = strconv.Itoa(remainder) + binaryStr
// 		dividend = quotient
// 	}
// 	return binaryStr, nil
// }

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "errors"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)
    binaryStr, err := decimalToBinary(n)
    if err != nil {
        fmt.Println(err)
    }
		printMaxConsecutiveOnes(binaryStr)
}

func printMaxConsecutiveOnes(binaryStr string) {
    max := 0
    currMax := 0
    for i := 0; i < len(binaryStr); i++ {
        if string(binaryStr[i]) == "0" {
            //we have finished counting  the 1'S
            currMax = 0
        } else {
            currMax++
            if currMax > max {
                max = currMax
            }
        }
    }
    fmt.Println(max)
}

func decimalToBinary(no int32) (string, error) {
    if no < 0 {
        return "", errors.New("input number must be a positive integer")
    }
    const divisor int32 = 2
    dividend := no
    var binaryStr string
    for dividend > 0 {
        quotient := dividend / divisor
        remainder := dividend % divisor
        binaryStr = strconv.Itoa(int(remainder)) + binaryStr
        dividend = quotient
    }
		return binaryStr, nil
}
func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

