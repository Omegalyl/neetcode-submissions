func maxArea(heights []int) int {
	mArea := 0
	currArea := 0
	i := 0
	n := len(heights) - 1

	for i < n {
		if heights[i] < heights[n] {
			currArea = (n - i) * heights[i]
			i++
		} else {
			currArea = (n - i) * heights[n]
			n--
		}

		if currArea > mArea {
			mArea = currArea
		}
	}
	return mArea
}
