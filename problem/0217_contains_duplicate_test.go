package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Input: nums = [1,2,3,1], Output: true",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: true,
		},
		{
			name: "Input: nums = [1,2,3,4], Output: false",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "Input: nums = [1,1,1,3,3,4,3,2,4,2], Output: true",
			args: args{
				nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, containsDuplicate(tt.args.nums), "containsDuplicate(%v)", tt.args.nums)
		})
	}
}
