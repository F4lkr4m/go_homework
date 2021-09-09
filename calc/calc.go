package calc

import (
	"errors"
	"strconv"
	"strings"
)

type Calc struct{}

func (calc *Calc) Solve(input string) (result float64, err error) {
	// catching of errors
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
			return
		}
	}()

	// clear input
	input = strings.Replace(input, " ", "", -1)
	// split input
	splittedInput := splitToOperands(input)
	// get polish expression
	polishExpression := convertToPolishSystem(splittedInput)
	// calculate
	return countPolishSystem(polishExpression), nil
}

// split string to []string with needed math operands and actions
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
		if input[i] == ',' {
			output += string('.')
			continue
		}

		output += string(input[i])
	}

	return output, len(output)
}

func reverseSlice(input []string) []string {
	for left, right := 0, len(input)-1; left < right; left, right = left+1, right-1 {
		input[left], input[right] = input[right], input[left]
	}
	return input
}

// converting to polish system func
func convertToPolishSystem(input []string) (output []string) {
	// stack for math operators
	operatorStack := make([]string, 0)

	// priorities of math operators
	var priority = map[string]int{"*": 2, "/": 2, "-": 1, "+": 1, "(": 0, ")": 0}

	// get every part of input and analyze it
	for i := 0; i < len(input); i++ {
		sign := input[i]
		if strings.ContainsAny("+-*/", sign) {
			// drop signs to output until operator with smaller be on the top of stack
			j := len(operatorStack) - 1
			for ; j >= 0 && (priority[operatorStack[j]] >= priority[sign]); j-- {
				var dropSign string
				operatorStack, dropSign = operatorStack[:len(operatorStack)-1], operatorStack[len(operatorStack)-1]
				output = append(output, dropSign)
			}
			// add operand to stack
			operatorStack = append(operatorStack, sign)
		} else if string(sign) == "(" {
			operatorStack = append(operatorStack, sign)
		} else if string(sign) == ")" {
			j := len(operatorStack) - 1
			// drop signs until (
			for ; j >= 0 && operatorStack[j] != "("; j-- {
				var dropSign string
				operatorStack, dropSign = operatorStack[:len(operatorStack)-1], operatorStack[len(operatorStack)-1]
				output = append(output, dropSign)
			}

			if len(operatorStack) == 0 {
				panic("Parsing expression error")
			} else {
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
		} else {
			output = append(output, sign)
		}
	}
	// add other operators in output
	operatorStack = reverseSlice(operatorStack)
	output = append(output, operatorStack...)

	return output
}

// counting func
func countPolishSystem(input []string) (result float64) {
	numberStack := make([]float64, 0)

	for i := 0; i < len(input); i++ {
		sign := input[i]

		if strings.ContainsAny("+-/*", sign) {
			var val float64
			// pop number in val and update stack
			numberStack, val = numberStack[:len(numberStack)-1], numberStack[len(numberStack)-1]
			// analyze sign
			switch sign {
			case "+":
				if len(numberStack) == 0 {
					numberStack = append(numberStack, val)
					continue
				}

				numberStack[len(numberStack)-1] += val
			case "-":
				if len(numberStack) == 0 {
					numberStack = append(numberStack, -val)
					continue
				}

				numberStack[len(numberStack)-1] -= val
			case "*":
				if len(numberStack) == 0 {
					panic("Parsing expression error")
				}

				numberStack[len(numberStack)-1] *= val
			case "/":
				if len(numberStack) == 0 {
					panic("Parsing expression error")
				}

				if val == 0 {
					panic("Zero division")
				}
				numberStack[len(numberStack)-1] = numberStack[len(numberStack)-1] / val
			}
		} else {
			// it is number -> add in number stack
			if s, err := strconv.ParseFloat(sign, 64); err == nil {
				numberStack = append(numberStack, s)
			} else {
				panic("Parsing expression error")
			}
		}
	}
	// result will be in first cell
	result = numberStack[0]
	return result
}
