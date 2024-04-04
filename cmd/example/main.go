package main

import (
	"flag"
	"fmt"
	lab2 "github.com/fujivara/lab2-arcprog"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Вираз для обробки")
	inputName       = flag.String("f", "", "Вхідний файл")
	outputName      = flag.String("o", "", "Вихідний файл")
)

func main() {
	flag.Parse()
	var readStream io.Reader
	var writeStream io.Writer

	if len(*inputExpression) < 1 && len(*inputName) < 1 {
		fmt.Fprintln(os.Stderr, "Вхідні дані мають бути надані")
		os.Exit(1)
	}

	if len(*inputExpression) > 0 {
		readStream = strings.NewReader(*inputExpression)
	} else if file, err := os.Open(*inputName); err == nil {
		defer file.Close()
		readStream = file
	} else {
		fmt.Fprintln(os.Stderr, "Вхідний файл не існує")
		os.Exit(1)
	}

	if len(*outputName) < 1 {
		writeStream = os.Stdout
	} else if file, err := os.Create(*outputName); err == nil {
		defer file.Close()
		writeStream = file
	} else {
		fmt.Fprintln(os.Stderr, "Помилка при відкритті вихідного файлу")
		os.Exit(1)
	}

	handler := lab2.ComputeHandler{R: readStream, W: writeStream}
	err := handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
