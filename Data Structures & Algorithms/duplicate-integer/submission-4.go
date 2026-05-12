func hasDuplicate(nums []int) bool {
    arr := make(map[int]struct{})
    for _, num := range nums {
        if _,ok := arr[num]; ok {
            return true
        }
        arr[num] = struct{}{}
    }
    return false
}
