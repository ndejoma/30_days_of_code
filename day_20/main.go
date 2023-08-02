// package main

// import "fmt"

// func main() {
// 	fmt.Println("using main")
// 	nums := []int{3, 5, 4, 7, 1}
// 	nums2 := []int{3, 5, 4, 7, 1}
// 	//3, 4, 5, 1, 7 //i = 0  - 4
// 	//
// 	result := bubbleSort(nums)
// 	fmt.Println(result)
// 	fmt.Println("**** THe better bubble sort ******")
// 	result2 := betterBubbleSort(nums2)
// 	fmt.Println(result2)
// }

// func bubbleSort(nums []int) []int {
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if nums[j] > nums[j+1] {
// 				//do the swap
// 				nums[j], nums[j+1] = nums[j+1], nums[j]
// 			}
// 		}
// 	}
// 	return nums
// }

//	func betterBubbleSort(nums []int) []int {
//		var isSorted bool
//		n := len(nums)
//		for !isSorted {
//			//in each iteration set the value of isSorted to true
//			isSorted = true
//			//loop over our list to see if it sorted
//			for i := 1; i < n; i++ {
//				//if the elements in the list are not sorted
//				if nums[i-1] > nums[i] {
//					nums[i-1], nums[i] = nums[i], nums[i-1]
//					isSorted = false
//				}
//			}
//		}
//		return nums
//	}
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func bubbleSort(nums []int32) {
	n := len(nums)
	swaps := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				//swap the the array elements
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swaps++
			}
		}
	}
	fmt.Printf("Array is sorted in %d swaps.\n", swaps)
	fmt.Printf("First Element: %d\n", nums[0])
	fmt.Printf("Last Element: %d\n", nums[n-1])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	// Write your code here
  bubbleSort(a)
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
