package main

import (
	"go_homework/calc"
)

func main() {
	calc := calc.Calc{}
	calc.Solve("(1,25 + 3) * 4 + 3")
	return
}
