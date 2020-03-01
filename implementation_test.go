package lab

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostfixToInfix(t *testing.T) {
	res, err := PostfixToInfix("4 2 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "((4-2)*3+5)", res)
	}

	res, err = PostfixToInfix("4 2 - -3 * 5 + 6 ^ 7 - 5 /")
	if assert.Nil(t, err) {
		assert.Equal(t, "(((4-2)*(-3)+5)^6-7)/5", res)
	}

	res, err = PostfixToInfix("4 2 -")
	if assert.Nil(t, err) {
		assert.Equal(t, "(4-2)", res)
	}

	res, err = PostfixToInfix("4 2 1 * 8 * + e f - g h * + * +")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("can't use an operator for less than two operands"), err)
	}

	res, err = PostfixToInfix("4 2 1 * 8 * + e f- g h * i + * +")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("add extra spaces between operands"), err)
	}

	res, err = PostfixToInfix("4 2 1 * 8 * + e f - g h * i + * +")
	if assert.Nil(t, err) {
		assert.Equal(t, "((4+2*1*8)+(e-f)*(g*h+i))", res)
	}

	res, err = PostfixToInfix("4 2 %")
	if assert.Error(t, err) {
		assert.Equal(t, errors.New("entered incorrect data"), err)
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("2 2 +")
	fmt.Println(res)
	// Output:
	// (2+2)
}
