func isValid(s string) bool {
	var stack []rune
	if len(s) == 1 {
		return false
	}
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
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
