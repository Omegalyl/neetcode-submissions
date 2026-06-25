func twoSum(numbers []int, target int) []int {
	i := 0
	n := len(numbers) - 1
	for{
		if i == n {
			break
		}

		if numbers[i] + numbers[n] > target {
			n--
			continue
		}

		if numbers[i] + numbers[n] == target {
			return []int{i+1, n+1}
		}

		if numbers[i] + numbers[n] < target {
			i++
			continue
		}
	}
	return nil
}
