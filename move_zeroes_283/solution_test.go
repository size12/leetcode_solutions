package move_zeroes_283

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"example 1",
			args{nums: []int{0, 1, 0, 3, 12}},
			[]int{1, 3, 12, 0, 0},
		},
		{
			"example 2",
			args{nums: []int{0}},
			[]int{0},
		},
		{
			"no zeroes",
			args{nums: []int{1, 2, 3, 4, 5}},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"only zeroes",
			args{nums: []int{0, 0, 0, 0, 0}},
			[]int{0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			assert.Equal(t, tt.want, tt.args.nums)
		})
	}
}
