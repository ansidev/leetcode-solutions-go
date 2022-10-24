package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: s = \"42\", Output: 42",
			args: args{
				s: "42",
			},
			want: 42,
		},
		{
			name: "Input: s = \"   -42\", Output: -42",
			args: args{
				s: "   -42",
			},
			want: -42,
		},
		{
			name: "Input: s = \"4193 with words\", Output: 4193",
			args: args{
				s: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "Input: s = \" \", Output: 0",
			args: args{
				s: " ",
			},
			want: 0,
		},
		{
			name: "Input: s = \"-\", Output: 0",
			args: args{
				s: "-",
			},
			want: 0,
		},
		{
			name: "Input: s = \"+\", Output: 0",
			args: args{
				s: "+",
			},
			want: 0,
		},
		{
			name: "Input: s = \"-+12\", Output: 0",
			args: args{
				s: "-+12",
			},
			want: 0,
		},
		{
			name: "Input: s = \"-91283472332\", Output: -2147483648",
			args: args{
				s: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "Input: s = \"9223372036854775808\", Output: 2147483647",
			args: args{
				s: "9223372036854775808",
			},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, myAtoi(tt.args.s), "myAtoi(%v)", tt.args.s)
		})
	}
}
