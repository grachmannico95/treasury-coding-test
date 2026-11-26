package question5

func GO(nums ...int) int {
	pairs := (len(nums) + 1) / 2

	res := make(chan int, pairs)

	for i := 0; i < len(nums); i += 2 {
		a := nums[i]
		b := a

		if i+1 < len(nums) {
			b = nums[i+1]
		}

		go func(a, b int) {
			res <- a * b
		}(a, b)
	}

	sum := 0
	for i := 0; i < pairs; i++ {
		sum += <-res
	}

	return sum
}
