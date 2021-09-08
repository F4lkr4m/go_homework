package main

import (
	"go_homework/uniq/utils"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test")

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	input := utils.Read(file)
	file1, _ := os.Create("outTest")

	utils.Write(file1, input)
	file1.Sync()
	file1.Close()
	return
}
