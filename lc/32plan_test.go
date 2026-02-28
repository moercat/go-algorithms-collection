package lc

import (
	"fmt"
	"testing"
)

func Test_longestValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses("))))())()()(()"))
}

func Test_longestValidParentheses2(t *testing.T) {
	fmt.Println(longestValidParentheses(")()())"))
}

func Test_longestValidParentheses3(t *testing.T) {
	fmt.Println(longestValidParentheses("(()"))
}
