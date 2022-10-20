package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Input: nums = [1,2,3,1], k = 3, Output: true",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			want: true,
		},
		{
			name: "Input: nums = [1,0,1,1], k = 1, Output: true",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			want: true,
		},
		{
			name: "Input: nums = [1,2,3,1,2,3], k = 2, Output: false",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, containsNearbyDuplicate(tt.args.nums, tt.args.k), "containsNearbyDuplicate(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
