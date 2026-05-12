func largestRectangleArea(heights []int) int {
    stack := []int{-1}
    heights = append(heights, 0)
    maxArea := 0

    for i := 0; i < len(heights); i++ {
        for len(stack) > 1 && heights[i] < heights[stack[len(stack) - 1]] {
            h := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            w := i - stack[len(stack)-1] - 1
            
            if h * w > maxArea {
                maxArea = h * w
            }
        }
        stack = append(stack, i)
    }

    return maxArea
}
