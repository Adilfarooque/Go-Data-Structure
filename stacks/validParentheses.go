package stacks

func IsValid(s string) bool {
	stack := []rune{} // Create an empty stack using a rune slice

	// Define a mapping of closing brackets to their corresponding opening brackets
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Iterate through the string
	for _, char := range s {
		if opening, ok := bracketMap[char]; ok {
			// If the current character is a closing bracket
			// Check if the stack is empty or if the top element of the stack matches the opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false // Invalid parentheses
			}
			stack = stack[:len(stack)-1] // Pop the opening bracket from the stack
		} else {
			stack = append(stack, char) // Push the character onto the stack
		}
	}

	// If the stack is empty at the end, all parentheses are valid
	return len(stack) == 0
}
