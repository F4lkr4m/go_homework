package uniq

import (
	"bufio"
	"go_homework/uniq/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func formatLineWithOptions(line string, opt utils.Options) (string, error) {
	if len(line) == 0 {
		return "", nil
	}
	formattedLine := line
	space, err := regexp.Compile(` +`)
	if err != nil {
		return "", err
	}
	formattedLine = space.ReplaceAllString(line, " ")

	if opt.FFields > 0 {
		separatedLine := strings.SplitN(formattedLine, " ", opt.FFields+1)
		if len(separatedLine) < opt.FFields+1 {
			return "", nil
		}
		formattedLine = " " + separatedLine[opt.FFields:][0]
	}

	if len(formattedLine) <= opt.SChars {
		return "", nil
	}
	formattedLine = formattedLine[opt.SChars:]

	if opt.CaseInsensitive {
		formattedLine = strings.ToLower(formattedLine)
	}

	return formattedLine, nil
}

type lineWithCounter struct {
	line          string
	formattedLine string
	counter       int
}

func Uniq(lines []string, opt utils.Options) (out []string, err error) {
	var countedLines []lineWithCounter

	formattedFirstLine, err := formatLineWithOptions(lines[0], opt)
	if err != nil {
		return out, err
	}
	countedLines = append(countedLines, lineWithCounter{lines[0], formattedFirstLine, 1})

	for _, line := range lines[1:] {
		formattedCurrentLine, err := formatLineWithOptions(line, opt)
		if err != nil {
			return out, err
		}
		if formattedCurrentLine == countedLines[len(countedLines)-1].formattedLine {
			countedLines[len(countedLines)-1].counter++
		} else {
			countedLines = append(countedLines, lineWithCounter{line, formattedCurrentLine, 1})
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
			return out, &ArgsError{}
		}
	}

	return out, nil
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

func writeResultData(data []string, outputFilename string) error {
	var outputWriter *bufio.Writer
	if outputFilename == "" {
		outputWriter = bufio.NewWriter(os.Stdout)
	} else {
		file, err := os.Create(outputFilename)
		if err != nil {
			return &OpenFileError{outputFilename}
		}
		defer file.Close()
		outputWriter = bufio.NewWriter(file)
	}

	err := utils.Write(outputWriter, data)
	if err != nil {
		return err
	}
	return nil
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
	result, err = Uniq(data, opt)
	if err != nil {
		return err
	}
	err = writeResultData(result, opt.OutputFilename)
	if err != nil {
		return err
	}

	return nil
}
