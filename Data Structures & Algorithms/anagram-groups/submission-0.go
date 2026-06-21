func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		sig := [26]int{}
		for _, c := range str {
			c = c - 'a'
			sig[c] += 1
		}
		if _, has := m[sig]; has {
			m[sig] = append(m[sig], str)
			continue
		}
		m[sig] = []string{str}
	}
	out := [][]string{}
	for _, v := range m  {
		out = append(out, v)
	}
	return out
}
