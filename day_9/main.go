package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n := 0
	fmt.Sscanf(scanner.Text(), "%d", &n)

	//the phoneBook map
	phoneBook := make(map[string]string)

	for i := 0; i < n; i++ {
		scanner.Scan()
		entry := strings.Split(scanner.Text(), " ")
		phoneBook[entry[0]] = entry[1]
	}
	for scanner.Scan() {
		name := scanner.Text()
		if phoneNo, ok := phoneBook[name]; ok {
			fmt.Printf("%s=%s\n", name, phoneNo)
		}else {
			fmt.Println("Not found")
		}
	}
}
