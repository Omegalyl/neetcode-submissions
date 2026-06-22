func topKFrequent(nums []int, k int) []int {
	idx := [2001][2]int{}
	for _, num := range nums {
		i := num + 1000
		if idx[i][0] == 0 {
			idx[i][1] = num
		}
		idx[i][0]++
	}
	sort.Slice(idx[:], func(i, j int) bool {return idx[i][0] > idx[j][0]})
	result := []int{}
	for _, v := range idx[:k] {
		result = append(result, v[1])
	}
	return result
}
