package calc

import (
	"strconv"
	"strings"
	"unicode"
)

func solve(input string) (float64, error) {
	if len(input) == 0 {
		return 0, &EmptyInputError{}
	}
	// clear input
	input = strings.ReplaceAll(input, " ", "")
	// split input
	splittedInput, err := splitToOperands(input)
	if err != nil {
		return 0, err
	}

	// get polish expression
	polishExpression, err := convertToPolishSystem(splittedInput)
	if err != nil {
		return 0, err
	}
	// calculate
	result, err := calculatePolishSystem(polishExpression)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// split string to []string with needed math operands and actions
func splitToOperands(input string) ([]string, error) {
	var output []string
	for i := 0; i < len(input); {
		sign, increment, err := getNextOperand(input, i)
		if err != nil {
			return output, err
		}
		output = append(output, sign)
		i += increment
	}

	return output, nil
}

// get next operand or action from input with the position
func getNextOperand(input string, position int) (output string, increment int, err error) {
	sign := input[position]

	// if sign is action
	if strings.ContainsAny("()+-/*", string(sign)) {
		output = output + string(sign)
		return output, 1, nil
	}

	// if sign is number
	for i := position; i < len(input) && !strings.ContainsAny("()+-/*", string(input[i])); i++ {
		if input[i] == ',' || input[i] == '.' {
			output += string('.')
			continue
		}

		if !unicode.IsDigit(rune(input[i])) {
			return output, len(output), &ParsingExprError{}
		}
		output += string(input[i])
	}

	return output, len(output), nil
}

func reverseSlice(input []string) []string {
	for left, right := 0, len(input)-1; left < right; left, right = left+1, right-1 {
		input[left], input[right] = input[right], input[left]
	}
	return input
}

// converting to polish system func
func convertToPolishSystem(input []string) ([]string, error) {
	// stack for math operators
	operatorStack := make([]string, 0)

	// priorities of math operators
	var priority = map[string]int{"*": 2, "/": 2, "-": 1, "+": 1, "(": 0, ")": 0}

	// get every part of input and analyze it
	var output []string
	for i := 0; i < len(input); i++ {
		sign := input[i]
		switch {
		case strings.ContainsAny("+-*/", sign):
			// drop signs to output until operator with smaller priority be on the top of stack
			j := len(operatorStack) - 1
			for ; j >= 0 && (priority[operatorStack[j]] >= priority[sign]); j-- {
				var dropSign string
				operatorStack, dropSign = operatorStack[:len(operatorStack)-1], operatorStack[len(operatorStack)-1]
				output = append(output, dropSign)
			}
			// add operand to stack
			operatorStack = append(operatorStack, sign)

		case string(sign) == "(":
			operatorStack = append(operatorStack, sign)
		case string(sign) == ")":
			// drop signs until (
			for j := len(operatorStack) - 1; j >= 0 && operatorStack[j] != "("; j-- {
				var dropSign string
				operatorStack, dropSign = operatorStack[:len(operatorStack)-1], operatorStack[len(operatorStack)-1]
				output = append(output, dropSign)
			}

			if len(operatorStack) == 0 {
				return output, &ParsingExprError{}
			}
			operatorStack = operatorStack[:len(operatorStack)-1]
		default:
			output = append(output, sign)
		}
	}
	// add other operators in output
	operatorStack = reverseSlice(operatorStack)

	output = append(output, operatorStack...)
	return output, nil
}

// counting func
func calculatePolishSystem(input []string) (float64, error) {
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
					return 0, &ParsingExprError{}
				}
				numberStack[len(numberStack)-1] *= val

			case "/":
				if len(numberStack) == 0 {
					return 0, &ParsingExprError{}
				}

				if val == 0 {
					return 0, &ZeroDivisionError{}
				}
				numberStack[len(numberStack)-1] = numberStack[len(numberStack)-1] / val
			}
			continue
		}
		// it is number -> add in number stack
		s, err := strconv.ParseFloat(sign, 64)
		if err != nil {
			return 0, &ParsingExprError{}
		}
		numberStack = append(numberStack, s)
	}
	// result will be in first cell
	return numberStack[0], nil
}
