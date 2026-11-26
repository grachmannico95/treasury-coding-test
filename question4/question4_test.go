package question4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGO(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want map[int]int
	}{
		{
			name: "digit repeats - from example",
			num:  1223334,
			want: map[int]int{
				1: 1,
				2: 2,
				3: 3,
				4: 1,
			},
		},
		{
			name: "single digit",
			num:  1,
			want: map[int]int{
				1: 1,
			},
		},
		{
			name: "zero",
			num:  0,
			want: map[int]int{
				0: 1,
			},
		},
		{
			name: "negative number",
			num:  -1223,
			want: map[int]int{
				1: 1,
				2: 2,
				3: 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GO(tc.num)
			assert.Equal(t, tc.want, got)
		})
	}
}
