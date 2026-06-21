func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for _, c := range s {
		if strings.Contains(t, string(c)) {
			t = strings.Replace(t, string(c), "", 1)
		}
	}
	return len(t) == 0
}
