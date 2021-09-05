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
	fmt.Println(splittedInput)
	polishExpression := convertToPolishSystem(splittedInput)
	fmt.Println(polishExpression)

	return calc.countPolishSystem(polishExpression), nil
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

func convertToPolishSystem(input []string) (output []string) {
	operatorStack := make([]string, 0)

	var priority = map[string]int{"*": 2, "/": 2, "-": 1, "+": 1, "(": 0, ")": 0}

	for i := 0; i < len(input); i++ {
		sign := input[i]
		if strings.ContainsAny("+-*/", string(sign)) {
			j := len(operatorStack) - 1
			for ; j >= 0 && (priority[operatorStack[j]] >= priority[sign]); j-- {
				var dropSign string
				operatorStack, dropSign = operatorStack[:len(operatorStack)-1], operatorStack[len(operatorStack)-1]
				output = append(output, dropSign)
			}
			operatorStack = append(operatorStack, sign)
		} else if string(sign) == "(" {
			operatorStack = append(operatorStack, sign)
		} else if string(sign) == ")" {
			j := len(operatorStack) - 1
			for ; j > -1 && operatorStack[j] != "("; j-- {
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
	operatorStack = reverseSlice(operatorStack)
	output = append(output, operatorStack...)

	return output
}

func (calc *Calc) countPolishSystem(input []string) (result float64) {
	numberStack := make([]float64, 0)

	for i := 0; i < len(input); i++ {
		sign := input[i]

		if strings.ContainsAny("()+-/*", string(sign)) {
			var val1 float64
			numberStack, val1 = numberStack[:len(numberStack)-1], numberStack[len(numberStack)-1]
			if len(numberStack) == 0 {
				panic("Parsing expression error")
			}
			switch sign {
			case "+":
				numberStack[len(numberStack)-1] = val1 + numberStack[len(numberStack)-1]
			case "-":
				numberStack[len(numberStack)-1] = numberStack[len(numberStack)-1] - val1
			case "*":
				numberStack[len(numberStack)-1] = val1 * numberStack[len(numberStack)-1]
			case "/":
				if val1 == 0 {
					panic("Zero division")
				}
				numberStack[len(numberStack)-1] = numberStack[len(numberStack)-1] / val1
			}
		} else {
			if s, err := strconv.ParseFloat(sign, 64); err == nil {
				numberStack = append(numberStack, s)
			} else {
				panic("Parsing expression error")
			}
		}
	}
	result = numberStack[0]
	return result
}
