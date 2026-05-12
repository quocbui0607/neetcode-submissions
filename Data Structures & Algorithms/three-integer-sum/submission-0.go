func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    hm := map[string]bool{}
    res := [][]int{}
    for i :=0;i<len(nums)-1;i++{
        val := nums[i]
        if val > 0 {
            break
        }
        l, r := i+1 , len(nums) - 1
        for l < r {
             threeSum := val + nums[l] + nums[r]
            if threeSum > 0 {
                r--
            } else if threeSum < 0 {
                l++
            } else {
                arr := []int{val, nums[l], nums[r]}
                if existed := hm[fmt.Sprintf("%v", arr)]; existed {
                    l++
                    r--
                    continue
                }
                res = append(res, arr)
                hm[fmt.Sprintf("%v", arr)] = true
                l++
                r--
            }
        }
    }

    return res
}  
