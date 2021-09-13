package uniq

import "fmt"

type OpenFileError struct {
	filepath string
}

func (e *OpenFileError) Error() string {
	return fmt.Sprintf("Can not open file with path: %s\n", e.filepath)
}
