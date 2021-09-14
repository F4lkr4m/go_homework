package calc

import "fmt"

type ParsingExprError struct{}

func (e *ParsingExprError) Error() string {
	return fmt.Sprintf("Parsing expression error\n")
}

type ZeroDivisionError struct{}

func (e *ZeroDivisionError) Error() string {
	return fmt.Sprintf("Zero division\n")
}

type EmptyInputError struct{}

func (e *EmptyInputError) Error() string {
	return fmt.Sprintf("Empty input\n")
}
