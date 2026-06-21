func isValid(s string) bool {
    pairs := map[rune]rune {
		')':'(',
		'}':'{',
		']':'[',
	}
	stack := []rune{}
	for _, ch := range s {
		pair , ok := pairs[ch]
		if ok {
			stackLen := len(stack)
			if stackLen == 0 {
				return false
			}
			if stack[len(stack) -1] == pair {
				stack = stack[:len(stack)-1]
				continue
			}

			return false
		}


		stack = append(stack, ch)
	}

	return len(stack) == 0
}
