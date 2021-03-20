package main

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for len(tokens) != 0 {
		top := tokens[0]
		tokens = tokens[1:]
		switch top {
		case "+", "-", "*", "/":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			switch top {
			case "+":
				stack[len(stack)-2] = a + b
			case "-":
				stack[len(stack)-2] = a - b
			case "*":
				stack[len(stack)-2] = a * b
			case "/":
				stack[len(stack)-2] = a / b
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, atoi(top))
		}
	}
	return stack[0]
}

func atoi(num string) int {
	var result int
	var neg bool
	if num[0] == '-' {
		neg = true
		num = num[1:]
	}
	for len(num) != 0 {
		result *= 10
		result += int(num[0] - '0')
		num = num[1:]
	}
	if neg {
		return -result
	} else {
		return result
	}
}
