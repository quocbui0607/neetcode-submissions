type Car struct {
    pos int
    time  float64
}

func carFleet(target int, position []int, speed []int) int {
    n := len(position)
    if n == 0 {
		return 0
	}

    cars := make([]Car, n)
    for i, _ := range cars {
        cars[i] = Car{
            pos: position[i],
            time: float64(target-position[i]) / float64(speed[i]),
        }
    }

    sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

    stack := []float64{}
    for _, car := range cars {
        if len(stack) == 0 || car.time > stack[len(stack)-1] {
			stack = append(stack, car.time)
		}
    }

    return len(stack)
}
