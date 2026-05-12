func characterReplacement(s string, k int) int {
    counts := make(map[byte]int)
    l := 0
    maxRepeat := 0
    res := 0
    
    for r := 0; r < len(s); r++ {
        counts[s[r]]++
        
        if counts[s[r]] > maxRepeat {
            maxRepeat = counts[s[r]]
        }

        for (r - l + 1) - maxRepeat > k {
            counts[s[l]]--
            l++
        }

        if (r - l + 1) > res {
            res = r - l + 1
        }
    }
    return res
}
