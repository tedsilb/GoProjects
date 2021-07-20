package calc_engine

import (
	"testing"
)

func TestGetNum(t *testing.T) {
	type test struct {
		inputExpr  string
		inputIndex int
		val        float64
		len        int
	}

	tests := []test{
		{inputExpr: "295xz", inputIndex: 0, val: float64(295), len: 3},
		{inputExpr: "hp295", inputIndex: 2, val: float64(295), len: 3},
		{inputExpr: "hp295xz", inputIndex: 2, val: float64(295), len: 3},
		{inputExpr: "295.44xz", inputIndex: 0, val: float64(295.44), len: 6},
		{inputExpr: "hp295.44", inputIndex: 2, val: float64(295.44), len: 6},
		{inputExpr: "hp295.44xz", inputIndex: 2, val: float64(295.44), len: 6},
	}

	for _, tc := range tests {
		gotVal, gotLen, err := getNum(&tc.inputExpr, tc.inputIndex)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if tc.val != gotVal {
			t.Errorf("want val %f, got %f", tc.val, gotVal)
		}
		if tc.len != gotLen {
			t.Errorf("want len %d, got %d", tc.len, gotLen)
		}
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  float64
	}{
		{name: "int add", input: "2+3+8+22", want: float64(35)},
		{name: "int sub", input: "10-5-15", want: float64(-10)},
		{name: "int mult", input: "3*5*3", want: float64(45)},
		{name: "int div", input: "20/4/2", want: float64(2.5)},
		{name: "int add, sub", input: "2+3-1", want: float64(4)},
		{name: "int add, sub, mult, div", input: "10-2*3+6/10", want: float64(3)},

		{name: "float add with spaces", input: "2.11+3.22+8+22", want: float64(35.33)},
		{name: "float sub", input: "10-5.7-15", want: float64(-10.7)},
		{name: "float mult", input: "3*5*3.2", want: float64(48)},
		{name: "float div", input: "20/4/2.5", want: float64(2)},
		{name: "fload add, sub", input: "2+3.8-1.3", want: float64(4.5)},
		{name: "float add, sub, mult, div with spaces", input: "10-2.3*3.4+6.82/3", want: float64(11)},

		{name: "int operation with spaces", input: " 10 - 2 * 3 + 6 / 10 ", want: float64(3)},
		{name: "float operation with spaces", input: " 10 - 2.3 * 3.4 + 6.82 / 3 ", want: float64(11)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Calculate(tc.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if got != tc.want {
				t.Errorf("Calculate() = %f, want %f", got, tc.want)
			}
		})
	}
}
