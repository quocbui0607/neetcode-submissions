func groupAnagrams(strs []string) [][]string {
    strMap := make(map[[26]int][]string)
    for _,str := range strs {
        var arrStr [26]int
        for _, c := range str {
            arrStr[int(c - 'a')] += 1
        }

        if _, ok := strMap[arrStr]; !ok {
            strMap[arrStr] = make([]string, 0)
        }
        strMap[arrStr] = append(strMap[arrStr], str)
    }
    
    res := make([][]string,0)
    for _,v := range strMap {
        res = append(res, v)
    }
    return res
}
