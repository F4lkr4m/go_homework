package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Calc struct {}

func (calc *Calc) Solve(input string) (result float64, err error) {
	// catching of errors
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
			return
		}
	}()

	input = strings.Replace(input, " ", "", -1)
	splittedInput := splitToOperands(input)
	for _, val := range splittedInput {

		if s, err := strconv.ParseFloat(val, 64); err == nil {
			fmt.Println(s) // 3.14159265
		} else {
			fmt.Println("Here we got an action")
		}
	}


	return
}

func splitToOperands(input string) (output []string) {
	for i := 0; i < len(input); {
		sign, increment := getNextOperand(input, i)
		output = append(output, sign)
		i += increment
	}



	return output
}

// get next operand or action from input with the position
func getNextOperand(input string, position int) (output string, increment int) {
	sign := input[position]

	// if sign is action
	if strings.ContainsAny("()+-/*", string(sign)) {
		output = output + string(sign)
		increment++
		return output, increment
	}

	// if sign is number
	for i := position; i < len(input) && !strings.ContainsAny("()+-/*", string(input[i])); i++ {
		output += string(input[i])
	}

	return output, len(output)
}

func convertToPolishSystem(input []string) (result []string, err error) {

	return
}