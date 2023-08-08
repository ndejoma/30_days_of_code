// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	r := scanner.Text()
// 	scanner.Scan()
// 	e := scanner.Text()
// 	expectedReturnDate := createDate(e)
// 	actualReturnDate := createDate(r)
// 	daysPassed := daysBetween(expectedReturnDate, actualReturnDate)
// 	var fine int
// 	if daysPassed <= 0 || sameDay(actualReturnDate, expectedReturnDate) {
// 		fine = 0
// 	} else {
// 		//calculate
// 		if sameMonth(actualReturnDate, expectedReturnDate) {
// 			fine = daysPassed * 15
// 		} else if sameYear(actualReturnDate, expectedReturnDate) {
// 			monthsPassed := monthsBetween(expectedReturnDate, actualReturnDate)
// 			fine = 500 * monthsPassed
// 		} else {
// 			fine = 1000
// 		}
// 	}
// 	fmt.Println(fine)
// }

// // check if two dates have the same day
// func sameDay(a, b time.Time) bool {
// 	return a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
// }

// // check if two date have the same month
// func sameMonth(a, b time.Time) bool {
// 	return a.Year() == b.Year() && a.Month() == b.Month()
// }

// // check if two date are within the same year
// func sameYear(a, b time.Time) bool {
// 	return a.Year() == b.Year()
// }

// // function to create a date
// func createDate(s string) time.Time {
// 	dates, err := splitAndParseToInt(s, " ")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return time.Date(dates[2], time.Month(dates[1]), dates[0], 0, 0, 0, 0, time.UTC)
// }

// // calculate the number of days between the return date and the expected day of return
// func daysBetween(e, r time.Time) int {
// 	duration := r.Sub(e)
// 	return int(duration.Hours() / 24)
// }

// // calculate the number of months passed
// //
// //	func monthsBetween(e, r time.Time) int {
// //		duration := r.Sub(e)
// //		return int()
// //	}
// func monthsBetween(e, r time.Time) int {
// 	years := r.Year() - e.Year()
// 	months := r.Month() - e.Month()

// 	return years*12 + int(months)
// }

// func splitAndParseToInt(s, sep string) ([]int, error) {
// 	parts := strings.Split(s, sep)
// 	ints := make([]int, len(parts))
// 	for i, part := range parts {
// 		n, err := strconv.Atoi(part)
// 		if err != nil {
// 			return nil, err
// 		}
// 		ints[i] = n
// 	}
// 	return ints, nil
// }


package main

import "fmt"

func main() {
    var d1, m1, y1 int
    var d2, m2, y2 int
    fmt.Scan(&d1, &m1, &y1)
    fmt.Scan(&d2, &m2, &y2)

    fine := 0

    if y1 > y2 {
        fine = 10000
    } else if y1 == y2 {
        if m1 > m2 {
            fine = 500 * (m1 - m2)
        } else if m1 == m2 && d1 > d2 {
            fine = 15 * (d1 - d2)
        }
    }

    fmt.Println(fine)
}
