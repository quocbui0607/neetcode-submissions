func findDuplicate(nums []int) int {
	for _, num := range nums {
		cur := int(math.Abs(float64(num)))
		if nums[cur] < 0 {
			return cur
		}
		nums[cur] *= -1
	}

	return -1
}
