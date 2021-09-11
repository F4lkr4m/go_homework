package main

import (
	"fmt"
	"go_homework/uniq"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	uniq.UniqManager()
	//calc.SolveArgsExpression()
	return
}
