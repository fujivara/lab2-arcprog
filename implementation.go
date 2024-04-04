package lab2

import (
	"fmt"
	"strings"
	"unicode"
)

func isOperator(op string) bool {
	switch op {
	case "+", "-", "*", "/", "^":
		return true
	default:
		return false
	}
}

func isValidOperand(operand string) bool {
	for _, char := range operand {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func isValidToken(token string) bool {
	return isOperator(token) || isValidOperand(token) || token == ","
}

func PrefixToInfix(prefix string) (string, error) {
	var stack []string
	tokens := strings.Fields(prefix)
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		if !isValidToken(token) {
			return "", fmt.Errorf("некоректний символ у виразі: %s", token)
		}
		if isOperator(token) {
			operands := make([]string, 0)
			for j := 0; j < 2; j++ {
				if len(stack) == 0 {
					return "", fmt.Errorf("недостатньо операндів для виразу")
				}
				operand := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				operands = append(operands, operand)
			}
			var infixExpr string
			if (token == "-" || token == "+") && i != 0 {
				infixExpr = fmt.Sprintf("(%s %s %s)", operands[0], token, operands[1])
			} else {
				infixExpr = fmt.Sprintf("%s %s %s", operands[0], token, operands[1])
			}
			stack = append(stack, infixExpr)
		} else {
			stack = append(stack, token)
		}
	}
	if len(stack) != 1 {
		return "", fmt.Errorf("недостатньо операторів для виразу")
	}
	return stack[0], nil
}
