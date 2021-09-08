package uniq

import "strings"

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

func FormatLinesWithOptions(lines []string, i bool, SChars int, FFields int) (out []string) {
	Outer:
		for _, line := range lines {
			formattedLine := line

			for i := 0; i < FFields; i++ {
				if index := strings.IndexByte(line, ' '); index == -1 {
					formattedLine = ""
					out = append(out, "")
					continue Outer
				} else {
					formattedLine = line[index:]
				}
			}

			if len(formattedLine) > SChars {
				formattedLine = formattedLine[SChars:]
			} else {
				formattedLine = ""
				out = append(out, "")
				continue Outer
			}

			if i {
				formattedLine = strings.ToLower(formattedLine)
			}

			out = append(out, formattedLine)
		}
	return out
}

func standardLogic(lines []string, formattedLines []string) (out []string){
	// just add the first one line
	out = append(out, lines[0])

	// check other lines
	for i := 1; i < len(lines) - 1; i++ {
		if formattedLines[i] != formattedLines[i - 1] {
			out = append(out, lines[i])
		}
	}

	if formattedLines[len(lines) - 1] != formattedLines[len(lines) - 2] {
		out = append(out, lines[len(lines) - 1])
	}

	return out
}

func Uniq(lines []string, opt Options) (out []string) {
	//originalLines := lines

	formattedLines := FormatLinesWithOptions(lines, opt.I, opt.SChars, opt.FFields)

	switch opt.WorkMode {
	case None:
		return standardLogic(lines, formattedLines)
	case C:

	case U:

	case D:

	}


	return
}