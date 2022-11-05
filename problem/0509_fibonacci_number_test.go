package problem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input n = 2, Output: 1",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "Input n = 3, Output: 2",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "Input n = 4, Output: 3",
			args: args{n: 4},
			want: 3,
		},
		{
			name: "Input n = 5, Output: 5",
			args: args{n: 5},
			want: 5,
		},
		{
			name: "Input n = 6, Output: 8",
			args: args{n: 6},
			want: 8,
		},
		{
			name: "Input n = 7, Output: 13",
			args: args{n: 7},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, fib(tt.args.n), "fib(%v)", tt.args.n)
		})
	}
}
