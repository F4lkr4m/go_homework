package calc

import (
	"flag"
	"fmt"
)

func getArg() (out string) {
	// get the expression
	flag.Parse()
	if len(flag.Args()) != 1 {
		return out
	}
	out = flag.Args()[0]
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
