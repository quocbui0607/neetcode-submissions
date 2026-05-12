func topKFrequent(nums []int, k int) []int {
    sort.Slice(nums, func (i,j int) bool {
        return nums[i] < nums[j]
    })

    numsMap := make([][]int, 0)
    lastVal := nums[0]
    count := 1
    for i:=1; i < len(nums);i++ {
        if lastVal == nums[i] {
            count++
            continue
        }
        numsMap = append(numsMap, []int{lastVal, count})
        lastVal = nums[i]
        count = 1
    }

    numsMap = append(numsMap, []int{lastVal, count})
    sort.Slice(numsMap, func (i,j int) bool {
        return numsMap[i][1] > numsMap[j][1]
    })

    res := make([]int,0)
    for _,val := range numsMap {
        if len(res) < k {res = append(res, val[0])}
    }
    return res
}
