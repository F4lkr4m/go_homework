package utils

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Read(buf *os.File) []string {
	scanner := bufio.NewScanner(buf)
	var result []string

	for scanner.Scan() {
		line := scanner.Text()

		result = append(result, line)
	}
	return result
}

func Write(buf *os.File, data []string) error {

	for _, item := range data {
		_, err := buf.WriteString(item + "\n")
		check(err)
	}
	return nil
}
