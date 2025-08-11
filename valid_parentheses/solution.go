package valid_parentheses

func isValid(s string) bool {
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	lastOpenRunes := make([]rune, 0, len(s)/2+1)

	for _, r := range s {
		if closeRune, ok := m[r]; ok {
			lastOpenRunes = append(lastOpenRunes, closeRune)
		} else {
			if len(lastOpenRunes) == 0 || lastOpenRunes[len(lastOpenRunes)-1] != r {
				return false
			}

			lastOpenRunes = lastOpenRunes[:len(lastOpenRunes)-1]
		}
	}

	return len(lastOpenRunes) == 0
}
