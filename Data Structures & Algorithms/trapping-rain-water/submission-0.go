func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }

    totalWaters := 0
    l, r := 0, len(height) - 1
    leftMax, rightMax := 0, 0

    for l < r {
        if height[l] < height[r] {
            if height[l] >= leftMax {
                leftMax = height[l]
            } else {
                totalWaters += leftMax - height[l]
            }
            l++
        } else {
            if height[r] >= rightMax {
                rightMax = height[r]
            } else {
                totalWaters += rightMax - height[r]
            }
            r--
        }  
    }
    return totalWaters
}
