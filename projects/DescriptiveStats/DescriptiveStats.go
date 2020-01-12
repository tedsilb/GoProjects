package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func calcMean(nums []int) float64 {
	var total int
	for _, value := range nums {
		total += value
	}
	return float64(total / len(nums))
}

func calcMedian(nums []int) float64 {
	// Sort the numbers
	sort.Ints(nums)

	if len(nums)%2 == 0 {
		// Even
		indexTargetLower := int((len(nums) - 1) / 2)
		indexTargetUpper := int((len(nums)-1)/2) + 1
		numLower := nums[indexTargetLower]
		numUpper := nums[indexTargetUpper]
		return float64((numUpper + numLower) / 2)
	}

	// Odd
	indexTarget := int(len(nums) / 2)
	return float64(nums[indexTarget])
}

func main() {
	userNums := gatherNumbersFromUser()

	mean := calcMean(userNums)
	fmt.Println("Mean:", mean)

	median := calcMedian(userNums)
	fmt.Println("Median:", median)
}
