func minEatingSpeed(piles []int, h int) int {
    l, r := 1.0, 1.0
    for _, p := range piles {
        if float64(p) > r {
            r = float64(p)
        }
    }

    for l < r {
        mid := math.Floor(l + (r - l)/2)
        spent := float64(0)
        for _, p := range piles {
            spent += math.Ceil(float64(p) / mid)
        }

        if spent <= float64(h) {
            r = mid

        } else {
            l = mid + 1

        }
    }
    return int(l)
}
