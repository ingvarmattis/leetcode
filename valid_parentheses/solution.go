package valid_parentheses

func isValid(s string) bool {
	const (
		circleOpen  rune = '('
		circleClose rune = ')'

		squareOpen  rune = '['
		squareClose rune = ']'

		figureOpen  rune = '{'
		figureClose rune = '}'
	)

	circleCount, squareCount, figureCount := 0, 0, 0
	lastOpenRunes := make([]rune, 0, len(s)/2+1)

	for _, r := range s {
		switch {
		case r == circleClose:
			if circleCount == 0 {
				return false
			}
			if lastOpenRunes[len(lastOpenRunes)-1] == circleOpen {
				circleCount--
				lastOpenRunes = lastOpenRunes[:len(lastOpenRunes)-1]
			} else {
				return false
			}
		case r == squareClose:
			if squareCount == 0 {
				return false
			}
			if lastOpenRunes[len(lastOpenRunes)-1] == squareOpen {
				squareCount--
				lastOpenRunes = lastOpenRunes[:len(lastOpenRunes)-1]
			} else {
				return false
			}
		case r == figureClose:
			if figureCount == 0 {
				return false
			}
			if lastOpenRunes[len(lastOpenRunes)-1] == figureOpen {
				figureCount--
				lastOpenRunes = lastOpenRunes[:len(lastOpenRunes)-1]
			} else {
				return false
			}
		case r == circleOpen:
			circleCount++
			lastOpenRunes = append(lastOpenRunes, r)
		case r == squareOpen:
			squareCount++
			lastOpenRunes = append(lastOpenRunes, r)
		case r == figureOpen:
			figureCount++
			lastOpenRunes = append(lastOpenRunes, r)
		}
	}

	return circleCount == 0 && squareCount == 0 && figureCount == 0
}
