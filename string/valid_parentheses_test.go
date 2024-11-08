package string

import (
	"fmt"
	"testing"
)

// 括号匹配问题
//Input: "()[]{}"
//Output: true

//Input: "([)]"
//Output: false

// go语言栈的设计
// 遇到左括号就进栈push，遇到右括号并且栈顶为与之对应的左括号，就把栈顶元素出栈。最后看栈里面还有没有 其他元素，如果为空，即匹配。
// 先进后出原则
func isValidParentheses(s string) bool { // 空字符串直接返回 true
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
		} else if ((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func isValidParentheses1(s string) bool {
	if len(s) == 0 {
		return false
	}
	stack := make([]rune, len(s))

	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			stack = append(stack, v)
		} else if (v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') ||
			(v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			(v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return true

}

func TestIsValidParentheses(t *testing.T) {
	str := "()[]{}"
	//str := "([)]"
	//str := "({[{}]})"
	fmt.Println(isValidParentheses(str))
	fmt.Println(isValidParentheses1(str))

}
