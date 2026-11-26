package question3

import "slices"

func GO(numsA, numsB []int) []int {
	if len(numsA) == 0 || len(numsB) == 0 {
		return []int{}
	}

	res := []int{}
	seen := make(map[int]bool)

	for _, v := range numsA {
		seen[v] = true
	}

	for _, v := range numsB {
		if seen[v] {
			res = append(res, v)
			delete(seen, v)
		}
	}

	slices.Sort(res)
	slices.Reverse(res)

	return res
}
