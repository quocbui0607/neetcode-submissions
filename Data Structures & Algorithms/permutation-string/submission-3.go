func checkInclusion(s1 string, s2 string) bool {
    if len(s2) < len(s1) {
        return false
    }

    s1Arr, s2Arr := [26]int{},  [26]int{}
    for i,c := range s1 {
        s1Arr[int(c - 'a')]++
        s2Arr[int(s2[i] - 'a')]++
    }

    if s1Arr == s2Arr {
        return true
    }
    for r := len(s1); r< len(s2);r++{
        s2Arr[int(s2[r] - 'a')]++
        s2Arr[int(s2[r-len(s1)] - 'a')]--

       if s1Arr == s2Arr {
        return true
        }
    }

    return false
}
