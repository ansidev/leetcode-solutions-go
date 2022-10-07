package problem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Input: digits = [1,2,3], Output: [1,2,4]",
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "Input: digits = [4,3,2,1], Output: [4,3,2,2]",
			args: args{
				digits: []int{4, 3, 2, 1},
			},
			want: []int{4, 3, 2, 2},
		},
		{
			name: "Input: digits = [9], Output: [1,0]",
			args: args{
				digits: []int{9},
			},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, plusOne(tt.args.digits), "plusOne(%v)", tt.args.digits)
		})
	}
}
