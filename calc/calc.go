package calc

import (
	"strings"
)

type Calc struct {}

func (calc *Calc) Solve(input string) (result float64, outErr error) {
	input = strings.Replace(input, " ", "", -1)


	defer func() {
		if r := recover(); r != nil {
			outErr = r.(error)
		}
	}()
	return
}

func splitToOperands(input string) (output []string, err error) {
	for i := 0; i < len(input); {

	}
	return 
}

func getNextOperand(input string, position int) (output string, nextSignIndex int, err error) {
	if position >= len(input) {
		panic("Out of index of input line")
	}

	return
}

func convertToPolishSystem(input []string) (result []string, err error) {

	return
}