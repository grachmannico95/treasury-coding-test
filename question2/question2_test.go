package question2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGO(t *testing.T) {
	testCases := []struct {
		name    string
		exclude []int
		nums    []int
		want    int
		wantErr string
	}{
		{
			name:    "all digit excluded - from example",
			exclude: []int{0, 1, 0},
			nums:    []int{1000, 1010},
			want:    0,
			wantErr: "not found",
		},
		{
			name:    "join, reverse, exclude - from example",
			exclude: []int{0, 1, 0},
			nums:    []int{1259, 2601},
			want:    62952,
		},
		{
			name:    "no exclude",
			exclude: nil,
			nums:    []int{1259, 2601},
			want:    10629521,
		},
		{
			name:    "one exclude",
			exclude: []int{9},
			nums:    []int{1259, 2601},
			want:    1062521,
		},
		{
			name:    "exclude number not present in slice",
			exclude: []int{7},
			nums:    []int{1259, 2601},
			want:    10629521,
		},
		{
			name:    "nums is nil",
			exclude: []int{0, 1, 0},
			want:    0,
			wantErr: "not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GO(tc.exclude, tc.nums...)
			assert.Equal(t, tc.want, got)
			if tc.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, tc.wantErr)
			}
		})
	}
}
