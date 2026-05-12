type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    encoded := make([]string, 0)
 for _, str := range strs {
    encoded = append(encoded, strconv.Itoa(len(str)))
    encoded = append(encoded, "#")
    encoded = append(encoded, str)
 }
 fmt.Println(encoded)
 return strings.Join(encoded,"")
}

func (s *Solution) Decode(encoded string) []string {
    res := make([]string, 0)
    i := 0
    for i < len(encoded) {
        j := i

       for encoded[j] != '#' {
            j++
        }

        l,_ := strconv.Atoi(encoded[i:j])        
        str := encoded[j+1: j+1+l]
         fmt.Println(str, l)

        res = append(res, str)
        i = j+1+l
    }
    return res
}
