func maxArea(heights []int) int {
    maxArea := 0
    l, r := 0, len(heights)-1
    for l < r {
    currVal := heights[l]

        distance := r - l
        if heights[l] > heights[r] {
            currVal = heights[r]
        }

        newArea := distance * currVal
        if newArea > maxArea {
            maxArea = newArea
        }

        if heights[l] <= heights[r] {
            l++
        } else {
            r--
        }
    }

    return maxArea
}
