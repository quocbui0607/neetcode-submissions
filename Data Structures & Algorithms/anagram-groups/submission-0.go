func groupAnagrams(strs []string) [][]string {
    strMap := make(map[string][]string)
    for _,str := range strs {
        runes := []rune(str)
        sort.Slice(runes, func (i, j int) bool {
            return runes[i] < runes[j]
        })
        runesConv := string(runes)
    fmt.Println(runesConv)

        if _, ok := strMap[runesConv]; !ok {
            strMap[runesConv] = make([]string, 0)
        }
        strMap[runesConv] = append(strMap[runesConv], str)
    }
    
    res := make([][]string,0)
    for _,v := range strMap {
        res = append(res, v)
    }
    return res
}
