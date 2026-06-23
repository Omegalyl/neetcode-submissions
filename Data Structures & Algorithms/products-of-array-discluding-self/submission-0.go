func productExceptSelf(nums []int) []int {
	// m := map[string]int{}
	result := []int{}

	for i := range nums {
		v := 1
		for j := range nums {
			if i == j {
				continue
			}
			v *= nums[j]
		}
		result = append(result, v)
	}
	return result
}
