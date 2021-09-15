package uniq

import (
	"bufio"
	"fmt"
	"go_homework/uniq/utils"
	"os"
	"strconv"
	"strings"
)

func formatLineWithOptions(line string, opt utils.Options) string {
	formattedLine := line

	if opt.FFields > 0 {
		separatedLine := strings.SplitN(line, " ", opt.FFields+1)
		fmt.Println(separatedLine)
		if len(separatedLine) < opt.FFields+1 {
			return ""
		}
		formattedLine = " " + strings.Join(separatedLine[opt.FFields:], " ")
	}

	if len(formattedLine) <= opt.SChars {
		return ""
	}
	formattedLine = formattedLine[opt.SChars:]

	if opt.CaseInsensitive {
		formattedLine = strings.ToLower(formattedLine)
	}

	return formattedLine
}

type lineWithCounter struct {
	line          string
	formattedLine string
	counter       int
}

func Uniq(lines []string, opt utils.Options) (out []string) {
	var countedLines []lineWithCounter

	countedLines = append(countedLines, lineWithCounter{lines[0], formatLineWithOptions(lines[0], opt), 1})

	for _, line := range lines[1:] {
		if formatLineWithOptions(line, opt) == countedLines[len(countedLines)-1].formattedLine {
			countedLines[len(countedLines)-1].counter++
		} else {
			countedLines = append(countedLines, lineWithCounter{line, formatLineWithOptions(line, opt), 1})
		}
	}

	for _, item := range countedLines {
		switch opt.WorkMode {
		case utils.None:
			out = append(out, item.line)
		case utils.D:
			if item.counter > 1 {
				out = append(out, item.line)
			}
		case utils.C:
			out = append(out, strconv.Itoa(item.counter)+" "+item.line)
		case utils.U:
			if item.counter == 1 {
				out = append(out, item.line)
			}
		case utils.Empty:
			return
		}
	}

	return
}

func getInputData(inputFilename string) (data []string, err error) {
	var inputReader *bufio.Reader
	if inputFilename == "" {
		inputReader = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(inputFilename)
		if err != nil {
			return data, &OpenFileError{inputFilename}
		}
		defer file.Close()
		inputReader = bufio.NewReader(file)
	}

	return utils.Read(inputReader), nil
}

func UniqManager() error {
	opt, err := utils.GetOptions()
	if err != nil {
		return err
	}

	var result []string

	data, err := getInputData(opt.InputFilename)
	if err != nil {
		return err
	}

	// if no data, just return
	if len(data) == 0 {
		return nil
	}
	result = Uniq(data, opt)

	var outputWriter *bufio.Writer
	if opt.OutputFilename == "" {
		outputWriter = bufio.NewWriter(os.Stdout)
	} else {
		file, err := os.Create(opt.OutputFilename)
		if err != nil {
			return &OpenFileError{opt.OutputFilename}
		}
		defer file.Close()
		outputWriter = bufio.NewWriter(file)
	}

	err = utils.Write(outputWriter, result)
	if err != nil {
		return err
	}
	return nil
}
