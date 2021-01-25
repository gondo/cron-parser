package main

import (
	"errors"
	"fmt"
	"github.com/gondo/cron-parser/internal/output"
	"github.com/gondo/cron-parser/internal/parser"
	"os"
)

func main() {
	input, err := processInput()
	checkError(err)

	cron, command, err := parser.Parse(input, parser.Slots)
	checkError(err)

	fmt.Println(output.Table(cron))
	fmt.Println(output.Row("command", command))
}

func processInput() (string, error) {
	args := os.Args[1:]
	if len(args) != 1 {
		return "", errors.New("invalid number of arguments")
	}
	return args[0], nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %v\n", err)
		os.Exit(1)
	}
}
