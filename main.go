package main

import (
	"go_homework/uniq"
)

func main() {

	err := uniq.UniqManager()
	if err != nil {
		panic(err)
	}
	return
}
