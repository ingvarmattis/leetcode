package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	div := 1
	for x/div >= 10 {
		div *= 10
	}

	for x != 0 {
		first := x / div
		last := x % 10

		if first != last {
			return false
		}

		x = (x % div) / 10
		div /= 100
	}

	return true
}
