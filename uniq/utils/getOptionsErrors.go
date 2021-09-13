package utils

import "fmt"

type TooManyArgsError struct{}

func (e *TooManyArgsError) Error() string {
	return fmt.Sprintf("Too many args error\n")
}
