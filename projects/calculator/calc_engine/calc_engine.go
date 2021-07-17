package calc_engine

import (
	"fmt"
	"strconv"
	"strings"
)

type Operation int

const (
	NO_OP Operation = iota
	ADD
	SUBTRACT
	MULTIPLY
	DIVIDE
)

var OPERATION_FROM_SYMBOL = map[rune]Operation{
	' ': NO_OP,
	'+': ADD,
	'-': SUBTRACT,
	'*': MULTIPLY,
	'/': DIVIDE,
}

// In a string starting at index i, get the full number represented.
// Returns the value, the parsed length, and any errors.
func getNum(expr *string, i int) (float64, int, error) {
	start := i
	end := start

	for {
		next := end + 1

		if next >= len(*expr) {
			break
		}

		nextChar := (*expr)[next]
		switch nextChar {
		case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			end = next
			continue
		}

		break
	}

	num, err := strconv.ParseFloat((*expr)[start:end+1], 64)

	return num, end - start + 1, err
}

// Perform a calculation of the provided mathematical expression.
// Order of operations is left-to-right.
func Calculate(expr string) (float64, error) {
	result := float64(0)
	nextOp := ADD

	if strings.ContainsAny(expr, "()") {
		return result, fmt.Errorf("unsupported characters provided")
	}

	for i := 0; i < len(expr); {
		if op, hasOp := OPERATION_FROM_SYMBOL[rune(expr[i])]; hasOp {
			if op != NO_OP {
				nextOp = op
			}
			i++
			continue
		}

		nextNum, nextNumLen, err := getNum(&expr, i)
		if err != nil {
			return result, fmt.Errorf("unexpected error parsing number from '%s' at index %d: %v", expr, i, err)
		}

		switch nextOp {
		case ADD:
			result += nextNum
		case SUBTRACT:
			result -= nextNum
		case MULTIPLY:
			result *= nextNum
		case DIVIDE:
			result /= nextNum
		}

		i += nextNumLen
	}

	return result, nil
}
