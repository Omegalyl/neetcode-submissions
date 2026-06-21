func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		var sig [26]int
		for _, c := range str {
			sig[c - 'a']++
		}
		m[sig] = append(m[sig], str)
	}
	out := make([][]string, 0, len(m))
	for _, v := range m  {
		out = append(out, v)
	}
	return out
}
