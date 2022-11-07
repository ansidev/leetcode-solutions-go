package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Input: nums = [3,2,4], target = 6, Output: [0,1]",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "Input: nums = [3,3], target = 6, Output: [0,1]",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
		{
			name: "Input: nums = [0,4,3,0], target = 0, Output: [0,3]",
			args: args{
				nums:   []int{0, 4, 3, 0},
				target: 0,
			},
			want: []int{0, 3},
		},
		{
			name: "Input: nums = [-1,-2,-3,-4,-5], target = -8, Output: [0,3]",
			args: args{
				nums:   []int{-1, -2, -3, -4, -5},
				target: -8,
			},
			want: []int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			assert.ElementsMatchf(t, tt.want, got, "Not match")
		})
	}
}
