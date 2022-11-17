package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMoneyAmount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: n = 10, Output: 16",
			args: args{
				n: 10,
			},
			want: 16,
		},
		{
			name: "Input: n = 1, Output: 0",
			args: args{
				n: 1,
			},
			want: 0,
		},
		{
			name: "Input: n = 2, Output: 1",
			args: args{
				n: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getMoneyAmount(tt.args.n), "getMoneyAmount(%v)", tt.args.n)
		})
	}
}
