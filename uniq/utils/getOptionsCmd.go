package utils

import (
	"errors"
	"flag"
)

type Mode int

const (
	None Mode = iota
	C
	D
	U
	Empty
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

func getOptionalMode(c bool, d bool, u bool) Mode {
	if ((c && u) || (c && d)) == true {
		return Empty
	}

	if c {
		return C
	}

	if d {
		return D
	}

	if u {
		return U
	}
	return None
}

func GetOptions() Options {
	options := Options{}

	c := flag.Bool("c", false, "count the number of occurrences of lines in the input")
	d := flag.Bool("d", false, "print only those lines that were repeated int the input data")
	u := flag.Bool("u", false, "print only those lines that were repeated int the input data")

	// get sensitive flag
	flag.BoolVar(&options.I, "i", false, "make program no sensitive to case")

	// flags with number argument
	flag.IntVar(&options.FFields, "f", 0, "ignore the first [num] of fields in line")
	flag.IntVar(&options.SChars, "s", 0, "ignore the first [num] of chars in line")

	flag.Parse()

	options.WorkMode = getOptionalMode(*c, *d, *u)

	if len(flag.Args()) > FileNum {
		panic(errors.New("Too much arguments\n"))
	}

	// get filenames
	switch len(flag.Args()) {
	case 1:
		options.InputFilename = flag.Args()[0]
	case 2:
		options.InputFilename = flag.Args()[0]
		options.OutputFilename = flag.Args()[1]
	}

	return options
}