package question5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGO(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "even count - from example",
			nums: []int{1, 2, 3, 4},
			want: 14,
		},
		{
			name: "odd count - from example",
			nums: []int{1, 2, 3, 4, 5},
			want: 39,
		},
		{
			name: "one pair",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "single number, mutiplied itself",
			nums: []int{2},
			want: 4,
		},
		{
			name: "no number",
			nums: []int{},
			want: 0,
		},
		{
			name: "negative number",
			nums: []int{1, -2, 1, -2},
			want: -4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GO(tc.nums...)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestGO_ManyNumber(t *testing.T) {
	nums := make([]int, 1001)
	for i := range nums {
		nums[i] = 1
	}

	got := GO(nums...)
	assert.Equal(t, 501, got)
}
