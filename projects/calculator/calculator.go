package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tedsilb/GoProjects/projects/calculator/engine"
)

func main() {
	fmt.Println("Simple calculator :)")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Type something to calculate, or type q to quit: ")

		scanner.Scan()
		expr := scanner.Text()
		if expr == "q" {
			os.Exit(0)
		}

		result, err := engine.Calculate(expr)
		if err != nil {
			fmt.Printf("Calculation error: %v\n", err)
			continue
		}

		fmt.Printf("Result: %f\n", result)
	}
}
