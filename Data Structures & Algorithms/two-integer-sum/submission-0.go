func twoSum(nums []int, target int) []int {
    data := map[int]int{}
    for i := 0; i < len(nums); i++ {
        if _, ok := data[nums[i]]; ok {
           return []int{data[nums[i]],i}
        }
        data[target - nums[i]] = i
    }

    return []int{}
}
