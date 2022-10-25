package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4], Output: 6",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "Input: [1], Output: 1",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "Input: [-2, 1], Output: 1",
			args: args{
				nums: []int{-2, 1},
			},
			want: 1,
		},
		{
			name: "Input: [-2, -1], Output: -1",
			args: args{
				nums: []int{-2, -1},
			},
			want: -1,
		},
		{
			name: "Input: [-1, -2], Output: -1",
			args: args{
				nums: []int{-1, -2},
			},
			want: -1,
		},
		{
			name: "Input: [-1], Output: -1",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
		{
			name: "Input: [5,4,-1,7,8], Output: 23",
			args: args{
				nums: []int{5, 4, -1, 7, 8},
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSubArray(tt.args.nums), "maxSubArray(%v)", tt.args.nums)
		})
	}
}
