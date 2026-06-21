func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for j, n:= range nums {
		if i, has := m[target - n]; has {
			return []int{i, j}
		}
		m[n] = j
	}
	return nil
}
