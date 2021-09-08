package uniq

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

	WorkMode Mode
	SChars int
	FFields int

	InputFilename string
	OutputFilename string
}