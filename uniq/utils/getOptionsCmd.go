package utils

import (
	"errors"
	"flag"
	"go_homework/uniq"
)

func getOptionalMode(c bool, d bool, u bool) uniq.Mode {
	if ((c && d) || (d && u) || (c && u)) == true {
		panic(errors.New("Invalid mode type.\n"))
	}

	if c {
		return uniq.C
	}

	if d {
		return uniq.D
	}

	if u {
		return uniq.U
	}
	return uniq.None
}

func GetOptions() uniq.Options {
	options := uniq.Options{}

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

	if len(flag.Args()) > uniq.FileNum {
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