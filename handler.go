package lab2

import (
	"bytes"
	"io"
)

type ComputeHandler struct {
	R io.Reader
	W io.Writer
}

func (ch *ComputeHandler) Compute() error {
	var inputBuffer bytes.Buffer
	const bufferSize = 1024
	buffer := make([]byte, bufferSize)

	for {
		n, err := ch.R.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		inputBuffer.Write(buffer[:n])
		if err == io.EOF {
			break
		}
	}

	input := inputBuffer.String()
	result, err := PrefixToInfix(input)
	if err != nil {
		return err
	}

	_, err = ch.W.Write([]byte(result))
	return err
}
