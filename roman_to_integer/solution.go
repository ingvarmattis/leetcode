package roman_to_integer

func romanToInt(s string) int {
	var sum, temp, prevVal int

	for _, current := range s {
		if prevVal == 0 {
			prevVal = m[current]
			temp = prevVal
			continue
		}

		currentVal := m[current]

		switch {
		case currentVal < prevVal:
			sum += temp
			temp = currentVal
		case currentVal == prevVal:
			temp += currentVal
		case currentVal > prevVal:
			temp = currentVal - temp
		}

		prevVal = currentVal
	}

	return sum + temp
}

var m = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}
