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

func standardLogic(lines []string, opt Options) (out []string){
	// just add the first one line
	out = append(out, lines[0])

	// check other lines
	for i := 1; i < len(lines); i++ {
		if formatLineWithOptions(lines[i], opt) != formatLineWithOptions(lines[i - 1], opt)  {
			out = append(out, lines[i])
		}
	}

	return out
}

type lineWithCounter struct {
	line string
	number int

}

func countingLogic(lines []string, opt Options) (out []string) {

	// just add the first one line
	var countedLines []lineWithCounter

	countedLines = append(countedLines, lineWithCounter{lines[0], 1})

	// check other lines
	for i := 1; i < len(lines); i++ {
		if formatLineWithOptions(lines[i], opt) == formatLineWithOptions(countedLines[len(countedLines) - 1].line, opt) {
			countedLines[len(countedLines) - 1].number++
		} else {
			countedLines = append(countedLines, lineWithCounter{lines[i], 1})
		}
	}

	for _, item := range countedLines {
		out = append(out, strconv.Itoa(item.number) + " " + item.line )
	}

	return out
}

func notRepeatingLogic(lines []string, opt Options) (out []string) {
	// check the first one
	if formatLineWithOptions(lines[0], opt) != formatLineWithOptions(lines[1], opt) {
		out = append(out, lines[0])
	}
	// check the main body
	for i := 1; i < len(lines) - 1; i++ {
		if formatLineWithOptions(lines[i], opt) != formatLineWithOptions(lines[i - 1], opt) &&
			formatLineWithOptions(lines[i], opt) != formatLineWithOptions(lines[i + 1], opt) {
			out = append(out, lines[i])
		}
	}

	// check the last one line
	if formatLineWithOptions(lines[len(lines) - 1], opt) != formatLineWithOptions(lines[len(lines) - 1], opt) {
		out = append(out, lines[len(lines) - 1])
	}

	return
}

func Uniq(lines []string, opt Options) (out []string) {
	switch opt.WorkMode {
	case None:
		return standardLogic(lines, opt)
	case C:
		return countingLogic(lines, opt)
	case U:
		return notRepeatingLogic(lines, opt)
	case D:

	}


	return
}