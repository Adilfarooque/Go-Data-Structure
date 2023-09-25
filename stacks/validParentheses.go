package stacks

func IsValid(s string) bool {
	stack := []rune{} // Create an empty stack using a rune slice

	for _,char:=range s{
		if char == '('{
			stack = append(stack, char)
		}else if char == ')'{
			if len(stack) == 0{
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack)== 0
}
