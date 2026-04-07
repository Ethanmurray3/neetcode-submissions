func isValid(s string) bool {
	var stack []rune

	for _, char := range s {
		if char == '(' {
			stack = append(stack, '(')
		} else if char == '{' {
			stack = append(stack, '{')
		} else if char == '[' {
			stack = append(stack, '[')
		} else if char == ')'{
			if len(stack) == 0 {
				return false
			}
			n := len(stack)
			if stack[n-1] == '(' {
				stack = stack[:n-1]
			} else {
				return false
			}
		} else if char == '}'{
			if len(stack) == 0 {
				return false
			}
			n := len(stack)
			if stack[n-1] == '{' {
				stack = stack[:n-1]
			} else {
				return false
			}
		} else if char == ']'{
			if len(stack) == 0 {
				return false
			}
			n := len(stack)
			if stack[n-1] == '[' {
				stack = stack[:n-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
