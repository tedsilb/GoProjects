package main

import (
	"bufio"
	"fmt"
	"math"
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
		return float64(nums[indexTargetLower+1]+nums[indexTargetLower]) / 2.0
	}

	// Odd
	indexTarget := int(len(nums) / 2)
	return float64(nums[indexTarget])
}

func calcVariance(nums []int) float64 {
	mean := calcMean(nums)

	sqDiffs := make([]int, 0)
	for _, value := range nums {
		sqDiffs = append(sqDiffs, int(math.Pow(float64(value)-mean, 2.0)))
	}

	return calcMean(sqDiffs)
}

func calcStDev(nums []int) float64 {
	return math.Pow(calcVariance(nums), 0.5)
}

func calcMin(nums []int) int {
	min := math.MaxInt8
	for _, value := range nums {
		if value < min {
			min = value
		}
	}
	return min
}

func calcMax(nums []int) int {
	max := math.MinInt8
	for _, value := range nums {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	userNums := gatherNumbersFromUser()

	fmt.Println("Mean:", calcMean(userNums))
	fmt.Println("Median:", calcMedian(userNums))
	fmt.Println("Variance:", calcVariance(userNums))
	fmt.Println("Standard Deviation:", calcStDev(userNums))
	fmt.Println("Min:", calcMin(userNums))
	fmt.Println("Max:", calcMax(userNums))
}
