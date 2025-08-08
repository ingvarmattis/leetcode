package longest_common_prefix

type str struct {
	arr []byte
	i   int
}

func (s *str) next() rune {
	if s.i >= len(s.arr) {
		return -1
	}

	b := s.arr[s.i]
	var r rune
	var size int

	switch {
	case b&0b10000000 == 0:
		r = rune(b)
		size = 1
	case b&0b11100000 == 0b11000000:
		r = (rune(b&0x1F) << 6) |
			rune(s.arr[s.i+1]&0x3F)
		size = 2
	case b&0b11110000 == 0b11100000:
		r = (rune(b&0x0F) << 12) |
			(rune(s.arr[s.i+1]&0x3F) << 6) |
			rune(s.arr[s.i+2]&0x3F)
		size = 3
	case b&0b11111000 == 0b11110000:
		r = (rune(b&0x7F) << 18) |
			(rune(s.arr[s.i+1]&0x3F) << 12) |
			(rune(s.arr[s.i+2]&0x3F) << 6) |
			rune(s.arr[s.i+3]&0x3F)
		size = 4
	default:
		return -1
	}

	s.i += size
	return r
}

func longestCommonPrefix(strs []string) string {
	ss := make([]*str, len(strs))
	for i, s := range strs {
		ss[i] = &str{
			arr: []byte(s),
			i:   0,
		}
	}

	var (
		r       rune
		returnS string
	)
	for i := 0; i <= len(ss); i++ {
		if i == len(ss) {
			i = 0
			returnS += string(r)
		}

		if i == 0 {
			tempR := ss[i].next()
			if tempR == -1 {
				return returnS
			}

			r = tempR
			continue
		}

		tempR := ss[i].next()
		if tempR == -1 {
			break
		}

		if tempR != r {
			break
		}
	}

	return returnS
}
