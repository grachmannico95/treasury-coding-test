package question4

func GO(num int) map[int]int {
	if num < 0 {
		num = num * -1
	}

	res := make(map[int]int)

	if num == 0 {
		res[0] = 1
		return res
	}

	for num > 0 {
		digit := num % 10
		res[digit]++
		num /= 10
	}

	return res
}
