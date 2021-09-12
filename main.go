package main

import (
	"go_homework/uniq"
)

func main() {

	err := uniq.UniqManager()
	//err := calc.SolveArgsExpression()
	if err != nil {
		panic(err)
	}
	return
}
