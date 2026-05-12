func lengthOfLongestSubstring(s string) int {
    obj := map[byte]bool{}
    l := 0
    longest := 0
    r := 0
    for r < len(s) {
        for obj[s[r]] {
            delete(obj, s[l])
            l++
        }
        if r - l + 1 > longest {
            longest = r - l + 1
        }
        obj[s[r]] = true
        r++
    }
    return longest
}
