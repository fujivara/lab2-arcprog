package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) { TestingT(t) }

type ComputeHandlerSuite struct{}

var _ = Suite(&ComputeHandlerSuite{})

func (s *ComputeHandlerSuite) TestHandlerCorrectResult(c *C) {
	input := "+ 2 * 3 4"
	expectedOutput := "2 + 3 * 4"

	reader := strings.NewReader(input)
	var writer bytes.Buffer

	compHandler := ComputeHandler{
		R: reader,
		W: &writer,
	}

	err := compHandler.Compute()
	c.Assert(err, IsNil)

	output := writer.String()
	c.Assert(output, Equals, expectedOutput)
}

func (s *ComputeHandlerSuite) TestHandlerSyntaxError(c *C) {
	input := "+ 2 * 3 4 )"
	expectedError := "некоректний символ у виразі: )"

	reader := strings.NewReader(input)
	var writer bytes.Buffer

	handler := ComputeHandler{
		R: reader,
		W: &writer,
	}

	err := handler.Compute()
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, expectedError)
}
