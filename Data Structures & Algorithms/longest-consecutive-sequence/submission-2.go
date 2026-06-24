func longestConsecutive(nums []int) int {
	if len(nums) < 1 {
		return len(nums)
	}
	sort.Ints(nums)
	maxConLen := 1
	currConLen := 1
	for i := 0; i < len(nums) - 1; i++ {
		diff := nums[i] - nums[i+1]
		if diff == 0 {
			continue 
		}

		if diff == -1 {
			currConLen++
		}  else {
			if currConLen > maxConLen {
				maxConLen = currConLen
			}
			currConLen = 1
		}
	}
	if currConLen > maxConLen {
		return currConLen
	}
	return maxConLen
}
