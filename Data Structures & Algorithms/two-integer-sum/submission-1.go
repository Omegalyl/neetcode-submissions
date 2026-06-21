func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i:= range nums {
		diff := target - nums[i]
		if _, has := m[diff]; has {
			return []int{m[diff], i}
		}
		m[nums[i]] = i
	}
	return nil
}
