package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: nums = [1,5,0,3,5], Output: 3",
			args: args{
				nums: []int{1, 5, 0, 3, 5},
			},
			want: 3,
		},
		{
			name: "Input: nums = [0], Output: 0",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "Input: nums = [1], Output: 1",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "Input: nums = [1,2,3,5], Output: 4",
			args: args{
				nums: []int{1, 2, 3, 5},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumOperations(tt.args.nums), "minimumOperations(%v)", tt.args.nums)
		})
	}
}
