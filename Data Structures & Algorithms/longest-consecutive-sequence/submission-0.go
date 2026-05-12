func longestConsecutive(nums []int) int {
    numsMap := make(map[int]bool)
    for _, val := range nums {
        numsMap[val] = true
    }

    longest := 0
    for k := range numsMap {
        currNum := k
        streak := 1
        for existed := numsMap[currNum + 1]; existed; existed = numsMap[currNum + 1] {
            currNum++
            streak++
        }

        if streak > longest {
            longest = streak
        }
   }
   return longest
}
