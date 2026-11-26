package question1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGO(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "has positive and negative - from example",
			arr:  []int{-1, 1},
			want: 0,
		},
		{
			name: "negative - from example",
			arr:  []int{-1, -7, -5},
			want: -2,
		},
		{
			name: "positive - from example",
			arr:  []int{1, 2, 1, 6},
			want: 3,
		},
		{
			name: "negative with missing number",
			arr:  []int{-2, -3, -4},
			want: -1,
		},
		{
			name: "positive with missing number",
			arr:  []int{2, 3, 4},
			want: 1,
		},
		{
			name: "single negative number",
			arr:  []int{-1},
			want: -2,
		},
		{
			name: "single positive number",
			arr:  []int{1},
			want: 2,
		},
		{
			name: "empty slice",
			arr:  []int{},
			want: 0,
		},
		{
			name: "zero",
			arr:  []int{0},
			want: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GO(tc.arr)
			assert.Equal(t, tc.want, got)
		})
	}
}
