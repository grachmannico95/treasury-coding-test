package question1

func GO(arr []int) int {
	positive, negative := false, false
	numbers := make(map[int]bool)

	if len(arr) == 0 {
		return 0
	}

	for _, val := range arr {
		numbers[val] = true

		if val > 0 {
			positive = true
		} else if val < 0 {
			negative = true
		}
	}

	if positive == negative {
		return 0
	}

	start := 1
	if negative {
		start = -1
	}

	for numbers[start] {
		if negative {
			start--
		} else {
			start++
		}
	}

	return start
}
