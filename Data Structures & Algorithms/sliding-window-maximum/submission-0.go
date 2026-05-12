func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	deque := []int{}
	res := []int{}

	for r := 0; r < len(nums); r++ {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[r] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, r)

		if deque[0] <= r-k {
			deque = deque[1:]
		}

		if r >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}

	return res
}