func twoSum(numbers []int, target int) []int {
	for i := range numbers {
		for j := range numbers {
			if numbers[i] + numbers[j] == target {
				return []int{i+1,j+1}
			}
		}
	}
	return nil
}
