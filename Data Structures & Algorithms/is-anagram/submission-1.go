func isAnagram(s string, t string) bool {
	sArr := [26]int{}
	tArr := [26]int{}
	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		sArr[int(s[i]- 'a')]++
		tArr[int(t[i]- 'a')]++
	}

	if sArr == tArr {
		return true
	}
	return false
}
