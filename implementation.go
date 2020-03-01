package lab

import (
	"errors"
	"regexp"
	s "strings"
)

func alphanumeric(matchStr string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9-]*$")
	return re.MatchString(matchStr)
}

func toString(slice []string) string {
	return s.Join(slice, "")
}

func hasOperator(str string) bool {
	return s.ContainsAny("+-*/^", str)
}
func isOperator(str string) bool {
	return s.Contains("+-*/^", str)
}

//lastElem, popStack
func popBack(stack []string) ([]string, []string) {
	lastIndex := len(stack) - 1
	return stack[lastIndex:], stack[:lastIndex]
}

func getLastEl(stack []string) []string {
	lastElem, _ := popBack(stack)
	return lastElem
}

func priority(char string) int {
	if s.Contains(char, "+") || s.Contains(char, "-") {
		return 1
	} else {
		return 2
	}
}

func PostfixToInfix(input string) (string, error) {
	var stack []string
	expression := s.TrimSpace(input)
	charArr := s.Split(expression, " ")
	if len(charArr) == 1 {
		err := errors.New("add extra spaces between operands")
		return toString(stack), err
	}
	for _, str := range charArr {
		if isOperator(str) {
			if len(stack) < 2 {
				err := errors.New("can't use an operator for less than two operands")
				return toString(stack), err
			}
			var leftArg, rightArg []string
			rightArg, stack = popBack(stack)
			leftArg, stack = popBack(stack)
			if priority(str) == 1 {
				stack = append(stack, "("+toString(leftArg)+str+toString(rightArg)+")")
			} else {
				stack = append(stack, toString(leftArg)+str+toString(rightArg))
			}
		} else if alphanumeric(str) {
			if hasOperator(str) {
				if s.HasPrefix(str, "-") {
					stack = append(stack, "("+str+")")
				} else {
					err := errors.New("add extra spaces between operands")
					return toString(stack), err
				}
			} else {
				stack = append(stack, str)
			}

		} else {
			err := errors.New("entered incorrect data")
			return toString(stack), err
		}
	}
	return toString(getLastEl(stack)), nil
}
