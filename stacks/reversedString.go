package stacks

func ReversedString(s string) string {
	var stack []rune

	for _, char := range s {
		stack = append(stack, char)
	}

	var reversedString string

	for len(stack) > 0 {
		reversedString = reversedString + string(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return reversedString
}
