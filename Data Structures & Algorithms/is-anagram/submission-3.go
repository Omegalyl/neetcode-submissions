func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[rune]int{}
	for _, c := range s {
		m[c] += 1
	}
	for _, c := range t {
		if _, has := m[c]; !has {
			return false
		}
		m[c] -= 1
		if m[c] < 0 {
			return false
		}
	}
	return true
}
