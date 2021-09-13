package calc

import (
	"flag"
	"fmt"
	"log"
)

func getArg() (out string) {
	// get the expression
	flag.Parse()

	if len(flag.Args()) == 1 {
		out = flag.Args()[0]
	} else {
		log.Fatal("Error! Need one argument with expression in \"*expression*\"")
	}
	return
}

func SolveArgsExpression() error {
	result, err := solve(getArg())
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
