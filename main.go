package main

import (
	"go_homework/calc"
)

func main() {

	err := calc.SolveArgsExpression()
	if err != nil {
		panic(err)
	}
	return
}
