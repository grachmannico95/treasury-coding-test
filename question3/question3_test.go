package question3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGO(t *testing.T) {
	testCases := []struct {
		name  string
		numsA []int
		numsB []int
		want  []int
	}{
		{
			name:  "one common number - from example",
			numsA: []int{1, 2},
			numsB: []int{1, 3},
			want:  []int{1},
		},
		{
			name:  "remove duplicate, sorted descending - from example",
			numsA: []int{1, 2, 2},
			numsB: []int{1, 2, 4},
			want:  []int{2, 1},
		},
		{
			name:  "unordered input",
			numsA: []int{3, 9, 1, 7},
			numsB: []int{7, 1, 3, 5},
			want:  []int{7, 3, 1},
		},
		{
			name:  "duplicate number on nums B",
			numsA: []int{1, 2},
			numsB: []int{2, 2, 2},
			want:  []int{2},
		},
		{
			name:  "negative number",
			numsA: []int{-1, -5, 0},
			numsB: []int{-5, -1, 8},
			want:  []int{-1, -5},
		},
		{
			name:  "no common number",
			numsA: []int{1, 2},
			numsB: []int{3, 4},
			want:  []int{},
		},
		{
			name:  "numA nil",
			numsA: []int{},
			numsB: []int{3, 4},
			want:  []int{},
		},
		{
			name:  "numB nil",
			numsA: []int{1, 2},
			numsB: []int{},
			want:  []int{},
		},
		{
			name:  "numA numB nil",
			numsA: []int{},
			numsB: []int{},
			want:  []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GO(tc.numsA, tc.numsB)
			assert.Equal(t, tc.want, got)
		})
	}
}
