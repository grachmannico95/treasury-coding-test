package question2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func GO(exclude []int, nums ...int) (int, error) {
	skip := make(map[rune]struct{}, len(exclude))
	digitChars := []rune("0123456789")

	for _, e := range exclude {
		if e >= 0 && e <= 9 {
			skip[digitChars[e]] = struct{}{}
		}
	}

	var joined strings.Builder
	for _, n := range nums {
		if n < 0 {
			return 0, fmt.Errorf("negative number is not supported: %d", n)
		}
		joined.WriteString(strconv.Itoa(n))
	}
	s := []rune(joined.String())

	digits := make([]rune, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		if _, ok := skip[s[i]]; ok {
			continue
		}
		digits = append(digits, s[i])
	}

	if len(digits) == 0 {
		return 0, errors.New("not found")
	}

	result, err := strconv.Atoi(string(digits))
	if err != nil {
		return 0, fmt.Errorf("result does not fit into int: %w", err)
	}

	return result, nil
}
