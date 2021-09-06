package calc

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

var sumTests = []struct {
	in  string
	out float64
}{
	{"1 + 3", 4},
	{"1 + 7", 8},
	{"1.5 + 3.5", 5.0},
	{"1.4 + 2.3", 3.7},
	{"1000000 + 1", 1000001},
	{"+3 + 4", 7},
	{"+3,25 + 4", 7.25},
}

var subtractionTests = []struct {
	in  string
	out float64
}{
	{"1 - 3", -2},
	{"1 - 7", -6},
	{"3.5 - 1.5", 2.0},
	{"1.5 - 1.5", 0},
	{"1000000 - 1", 999999},
	{"- 3 - 1", -4},
}

var divTests = []struct {
	in  string
	out float64
}{
	{"4 / 2", 2},
	{"9 / 3", 3},
	{"10 / 5", 2.0},
	{"9.3 / 3", 3.1},
}

var mulTests = []struct {
	in  string
	out float64
}{
	{"1 * 3", 3},
	{"2 * 7", 14},
	{"1.5 * 3.5", 5.25},
	{"9999 * 999", 9989001},
}

var tests = [][]struct {
	in  string
	out float64
}{
	sumTests,
	subtractionTests,
	divTests,
	mulTests,
}

func TestElementaryOperations(t *testing.T) {
	calc := Calc{}
	for _, cases := range tests {
		for _, tt := range cases {
			t.Run(tt.in, func(t *testing.T) {
				val, err := calc.Solve(tt.in)
				if err != nil {
					t.Errorf("Something went wrong, error in solving elementary operations.")
				}
				tolerance := 1e-15
				if diff := math.Abs(val - tt.out); diff > tolerance {
					t.Errorf("Something went wrong, error in solving elementary operations\nExpected: %f\nGot: %f", tt.out, val)
				}
			})
		}
	}
}

var hardTests = []struct {
	in  string
	out float64
}{
	{"1 + 3 + (4 * (3 + (4 / 4)))", 20},
	{"1 - 4 * (3 + 4)", -27},
	{"2 * (3 + 4 * 4 * 5 * 5) * 3", 2418},
	{"3 + 2 * 4 * 2 / (1 - 3)", -5},
}

func TestHardOperations(t *testing.T) {
	calc := Calc{}
	for _, tt := range hardTests {
		t.Run(tt.in, func(t *testing.T) {
			val, err := calc.Solve(tt.in)
			if err != nil {
				t.Errorf("Something went wrong, error in solving elementary operations.")
			}
			require.Equal(t, tt.out, val, "Two numbers should be the same.")
		})
	}
}

var errorTests = []struct {
	in  string
	out string
}{
	{"1 + 3 + (4 * (3 + (4 / 1))", "Parsing expression error"},
	{"1.2 + 3 + (4 * (3 + (4 / 1))))", "Parsing expression error"},
	{"Hello there", "Parsing expression error"},
	{"* 3", "Parsing expression error"},
	{"/ 3", "Parsing expression error"},
	{"1 + 3 + (4 * (3 + (4 / 0)))", "Zero division"},
	{"1 + 3 + (4 * (3 + (4 / (1 - 1))))", "Zero division"},
}

func TestErrorInput(t *testing.T) {
	calc := Calc{}
	for _, tt := range errorTests {
		t.Run(tt.in, func(t *testing.T) {
			_, err := calc.Solve(tt.in)
			require.EqualError(t, err, tt.out)
		})
	}
}
