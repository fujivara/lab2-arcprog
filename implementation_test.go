package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	_ "strings"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type PrefixToInfixSuite struct{}

var _ = Suite(&PrefixToInfixSuite{})

func (s *PrefixToInfixSuite) TestPrefixToInfixEmptyInput(c *C) {
	infix, err := PrefixToInfix("")
	c.Check(err, NotNil)
	c.Check(infix, Equals, "")
}

func (s *PrefixToInfixSuite) TestPrefixToInfixInvalidCharacter(c *C) {
	invalidPrefix := "1 2 + #"
	infix, err := PrefixToInfix(invalidPrefix)
	c.Check(err, NotNil)
	c.Check(infix, Equals, "")
}

func (s *PrefixToInfixSuite) TestToInfixValidInputs(c *C) {
	tests := []struct {
		prefix   string
		expected string
	}{
		{"+ 5 * - 4 2 3", "5 + (4 - 2) * 3"},
		{"/ * - 20 4 + 8 2 5", "(20 - 4) * (8 + 2) / 5"},
		{"+ * 2 3 4", "2 * 3 + 4"},
		{"- * 3 + 2 5 / 12 6", "3 * (2 + 5) - 12 / 6"},
		{"* + 3 5 2", "(3 + 5) * 2"},
	}

	for _, test := range tests {
		infix, err := PrefixToInfix(test.prefix)
		if c.Check(err, IsNil) {
			c.Check(infix, Equals, test.expected)
		}
	}
}

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("* - 4 2 6")
	fmt.Println(res)

	// Output:
	// (4 - 2) * 6
}
