func isValid(s string) bool {
    stack := []rune{}
    pair := map[rune]rune {
        '}': '{',
        ']': '[',
        ')': '(',
    }
    for _, c := range s {
        if c == '{' || c == '[' || c == '(' {
            stack = append(stack,c)
        }

        if c == '}' || c == ']' || c == ')' {
            if len(stack) == 0 || pair[c] != stack[len(stack)-1] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}
