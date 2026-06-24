func isPalindrome(s string) bool {
	i := 0
	s = strings.ToLower(s)
	n := len(s) - 1
	for {
		if i >= n {
			return true
		}

		if !isAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isAlphanumeric(s[n]) {
			n--
			continue
		}
		if s[i] != s[n] {
			return false
		}
		n--
		i++
	}
	return true
}

func isAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}