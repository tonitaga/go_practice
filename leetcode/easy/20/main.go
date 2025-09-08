package main

// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/description/

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, ch := range s {
		len := len(stack)

		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
			continue
		}

		if len == 0 {
			return false
		}

		top := stack[len-1]
		if (ch == ')' && top == '(') || (ch == '}' && top == '{') || (ch == ']' && top == '[') {
			stack = stack[:len-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
