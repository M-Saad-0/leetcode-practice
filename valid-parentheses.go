package main

func isValid(s string) bool {
	var stk []byte = []byte{}
	// if s[0] == '}' || s[0] == ')' || s[0] == ']' {
	// 	return false
	// }
	// if s[len(s)-1] == '(' || s[len(s)-1] == '{' || s[len(s)-1] == '[' {
	// 	return false
	// }
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[byte]byte{')': '(', '}': '{', ']': '['}

	for i := 0; i < len(s); i++ {
		isOpening := s[i] == '(' || s[i] == '[' || s[i] == '{'

		if isOpening {
			stk = append(stk, s[i])
			continue
		}
		isClosing := s[i] == ')' || s[i] == '}' || s[i] == ']'

		if isClosing && len(stk) != 0 && stk[len(stk)-1] == pairs[s[i]] {

			stk = pop(stk)

		} else {
			return false
		}
	}

	return len(stk) == 0
}

func pop(slc []byte) []byte {
	if len(slc) == 0 {
		return []byte{}
	}
	return slc[:len(slc)-1]
}

// "{{("
