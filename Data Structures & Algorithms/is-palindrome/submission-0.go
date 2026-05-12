func isPalindrome(s string) bool {
    l, r := 0, len(s) -1 
    for l < r {
        		cl, cr := rune(s[l]), rune(s[r])

        if !unicode.IsLetter(cl) && !unicode.IsDigit(cl){
            l++
            continue
        }

         if !unicode.IsLetter(cr) && !unicode.IsDigit(cr){
            r--
            continue
        }

        if unicode.ToLower(cl) != unicode.ToLower(cr) {
            return false
        }

        l++
        r--
    }
    

    return true
}
