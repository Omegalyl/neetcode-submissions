func twoSum(numbers []int, target int) []int {
	i := 0
	n := len(numbers) - 1
	for i<n {
		if numbers[i] + numbers[n] > target {
			n--
		} else if numbers[i] + numbers[n] < target {
			i++
		} else {
			return []int{i+1, n+1}
		}
	}
	return nil
}
