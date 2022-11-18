package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: nums = [1,2,3,4], Output: 24",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 24,
		},
		{
			name: "Input: nums = [-4,-3,-2,-1,60], Output: 720",
			args: args{
				nums: []int{-4, -3, -2, -1, 60},
			},
			want: 720,
		},
		{
			name: "Input: nums = [-1,-2,1,2,3], Output: 6",
			args: args{
				nums: []int{-1, -2, 1, 2, 3},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumProduct(tt.args.nums), "maximumProduct(%v)", tt.args.nums)
		})
	}
}
