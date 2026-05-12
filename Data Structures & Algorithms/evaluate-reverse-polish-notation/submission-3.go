func evalRPN(tokens []string) int {
    stack := []int{}
    for _, c := range tokens {
        val, err := strconv.Atoi(c)
        if  err == nil {
            stack = append(stack,val)
        } else {
        last := stack[len(stack)-2]
        first := stack[len(stack)-1]
        stack = stack[:len(stack)-2]

        var res int
            switch c {
            case "+":
                res = last + first
            case "-":
                res = last - first
            case "*":
                res = last * first
            case "/":
                res = last / first
            }
            stack = append(stack, res)
        }
        
    }

    return int(stack[0])
}
