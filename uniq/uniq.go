package uniq

import (
	"strconv"
	"strings"
)

type Mode int

const (
	None Mode = iota
	C
	D
	U
)

const FileNum = 2

type Options struct {
	I bool  // no case-sensitive
	SChars int
	FFields int

	WorkMode Mode

	InputFilename string
	OutputFilename string
}

func formatLineWithOptions(line string, opt Options) (out string){
	formattedLine := line

	for i := 0; i < opt.FFields; i++ {
		if index := strings.IndexByte(line, ' '); index == -1 {
			out = ""
			return
		} else {
			formattedLine = line[index:]
		}
	}

	if len(formattedLine) > opt.SChars {
		formattedLine = formattedLine[opt.SChars:]
	} else {
		out = ""
		return
	}

	if opt.I {
		formattedLine = strings.ToLower(formattedLine)
	}

	out = formattedLine
	return
}

type lineWithCounter struct {
	line string
	number int

}

func Uniq(lines []string, opt Options) (out []string) {
	var countedLines []lineWithCounter

	countedLines = append(countedLines, lineWithCounter{lines[0], 1})

	for i := 1; i < len(lines); i++ {
		if formatLineWithOptions(lines[i], opt) == formatLineWithOptions(countedLines[len(countedLines) - 1].line, opt) {
			countedLines[len(countedLines) - 1].number++
		} else {
			countedLines = append(countedLines, lineWithCounter{lines[i], 1})
		}
	}

	for _, item := range countedLines {
		switch opt.WorkMode {
		case None:
			out = append(out, item.line)
		case D:
			if item.number > 1 {
				out = append(out, item.line)
			}
		case C:
			out = append(out, strconv.Itoa(item.number)+" "+item.line)
		case U:
			if item.number == 1 {
				out = append(out, item.line)
			}
		}
	}

	return
}