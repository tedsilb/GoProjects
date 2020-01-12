package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gatherNumbersFromUser() []int {
	// Create a dynamic array
	nums := make([]int, 0)

	// Create stdin scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Get numbers until we want to stop
	for {
		fmt.Print("Enter a whole number, or press Enter to stop: ")

		// Read a line from stdin
		scanner.Scan()
		text := scanner.Text()

		// Make sure we got input
		if len(text) != 0 {
			num, _ := strconv.Atoi(text)
			nums = append(nums, num)
		} else {
			break
		}
	}

	return nums
}

func main() {
	userNums := gatherNumbersFromUser()

	fmt.Println(userNums)
}
