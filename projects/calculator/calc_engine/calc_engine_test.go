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
	type test struct {
		input string
		want  float64
	}

	tests := []test{
		{input: "2+3+8+22", want: float64(35)},
		{input: " 10 - 5 - 15 ", want: float64(-10)},
		{input: "3*5*3", want: float64(45)},
		{input: "20/4/2", want: float64(2.5)},
		{input: "2+3-1", want: float64(4)},
		{input: "10-2*3+6/10", want: float64(3)},

		{input: "2.11 + 3.22 + 8 + 22", want: float64(35.33)},
		{input: "10-5.7-15", want: float64(-10.7)},
		{input: "3*5*3.2", want: float64(48)},
		{input: "20/4/2.5", want: float64(2)},
		{input: "2+3.8-1.3", want: float64(4.5)},
		{input: "10 - 2.3 * 3.4 + 6.82 / 3", want: float64(11)},
	}

	for _, tc := range tests {
		got, err := Calculate(tc.input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}
