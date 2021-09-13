package uniq

import (
	"bufio"
	"go_homework/uniq/utils"
	"os"
	"strconv"
	"strings"
)

func formatLineWithOptions(line string, opt utils.Options) string {
	formattedLine := line

	for i := 0; i < opt.FFields; i++ {
		index := strings.IndexByte(line, ' ')
		if index == -1 {
			return ""
		}
		formattedLine = line[index:]
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
	line    string
	counter int
}

func Uniq(lines []string, opt utils.Options) (out []string) {
	var countedLines []lineWithCounter

	countedLines = append(countedLines, lineWithCounter{lines[0], 1})

	//for i := 1; i < len(lines); i++ {
	for _, line := range lines[1:] {
		if formatLineWithOptions(line, opt) == formatLineWithOptions(countedLines[len(countedLines)-1].line, opt) {
			countedLines[len(countedLines)-1].counter++
		} else {
			countedLines = append(countedLines, lineWithCounter{line, 1})
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

func UniqManager() error {
	opt := utils.GetOptions()
	if opt == nil {
		return &GetOptionsError{}
	}

	var result []string
	var data []string
	var inputReader *bufio.Reader
	if opt.InputFilename == "" {
		inputReader = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(opt.InputFilename)
		if err != nil {
			return &OpenFileError{opt.InputFilename}
		}
		defer file.Close()
		inputReader = bufio.NewReader(file)
	}

	data = utils.Read(inputReader)
	// if no data, just return
	if len(data) == 0 {
		return nil
	}
	result = Uniq(data, *opt)

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

	err := utils.Write(outputWriter, result)
	if err != nil {
		return err
	}
	return nil
}
