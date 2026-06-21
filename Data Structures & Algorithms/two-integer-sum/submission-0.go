func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i:= range nums {
		diff := target - nums[i]
		if _, has := m[diff]; !has {
			m[nums[i]] = i
			continue
		}
		return []int{m[diff], i}
	}
	return []int{}
}
