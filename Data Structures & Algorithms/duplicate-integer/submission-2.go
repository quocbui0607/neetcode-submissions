func hasDuplicate(nums []int) bool {
    arr := map[int]bool{}
    for _, num := range nums {
        if (arr[num]) {
            return true
        }
        arr[num] = true
    }
    return false
}
