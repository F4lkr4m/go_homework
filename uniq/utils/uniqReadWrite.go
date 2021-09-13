package utils

import (
	"bufio"
)

func Read(buf *bufio.Reader) []string {
	scanner := bufio.NewScanner(buf)
	var result []string

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

func Write(buf *bufio.Writer, data []string) error {
	for _, item := range data {
		_, err := buf.WriteString(item + "\n")
		if err != nil {
			return err
		}
	}
	err := buf.Flush()
	if err != nil {
		return err
	}
	return nil
}
