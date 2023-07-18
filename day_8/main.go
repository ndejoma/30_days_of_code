package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// var n int = 5
	var arr []int32 = []int32{1, 2, 3, 4, 5}
	// for i := len(arr) - 1; i >= 0; i-- {
	// 	arr[]
	// 	fmt.Println(arr)
	// 	fmt.Println(arr[i], i)
	// }
	var arrStr string
	for i := len(arr) - 1; i >= 0; i-- {
		arrStr += fmt.Sprintf("%s ", strconv.Itoa(int(arr[i])))
	}
	arrStr = strings.TrimSpace(arrStr)
	fmt.Println(arrStr)
}
