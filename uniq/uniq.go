package uniq

type mode int

const (
	c mode = iota
	d
	u
)

type Options struct {
	i bool  // no case-sensitive

	workMode mode
	sChars int
	fFields int
}