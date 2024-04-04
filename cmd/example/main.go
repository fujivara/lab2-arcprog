package main

import (
	"flag"
	"fmt"
	lab2 "github.com/roman-mazur/architecture-lab-2"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputName       = flag.String("f", "", "Input file name")
	outputName      = flag.String("o", "", "Output file name")
)

func main() {
	flag.Parse()
	var readStream io.Reader
	var writeStream io.Writer

	if len(*inputExpression) < 1 && len(*inputName) < 1 {
		fmt.Fprintln(os.Stderr, "Input files must be provided")
		os.Exit(1)
	}

	if len(*inputExpression) > 0 {
		readStream = strings.NewReader(*inputExpression)
	} else if file, err := os.Open(*inputName); err == nil {
		defer file.Close()
		readStream = file
	} else {
		fmt.Fprintln(os.Stderr, "Input file does not exist")
		os.Exit(1)
	}

	if len(*outputName) < 1 {
		writeStream = os.Stdout
	} else if file, err := os.Create(*outputName); err == nil {
		defer file.Close()
		writeStream = file
	} else {
		fmt.Fprintln(os.Stderr, "Cannot write to output file")
		os.Exit(1)
	}

	handler := lab2.ComputeHandler{R: readStream, W: writeStream}
	err := handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
