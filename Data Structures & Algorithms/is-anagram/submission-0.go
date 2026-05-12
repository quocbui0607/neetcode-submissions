func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := map[rune]int{}
	for _, c := range s {
		if _, ok := charMap[c]; !ok {
			charMap[c] = 0
		}
		charMap[c] += 1
	}

	for _,c := range t {
		if _,ok := charMap[c]; !ok || charMap[c] == 0 {
			return false			
		}
		charMap[c] -= 1
	}
	return true
}
