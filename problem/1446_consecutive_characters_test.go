package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxPower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: s = \"leetcode\", Output: 2",
			args: args{
				s: "leetcode",
			},
			want: 2,
		},
		{
			name: "Input: s = \"abbcccddddeeeeedcba\", Output: 5",
			args: args{
				s: "abbcccddddeeeeedcba",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxPower(tt.args.s), "maxPower(%v)", tt.args.s)
		})
	}
}
