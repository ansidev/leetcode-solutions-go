package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_guessNumber(t *testing.T) {
	type args struct {
		n    int
		pick int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: n = 1, pick = 1, Output: 1",
			args: args{
				n:    1,
				pick: 1,
			},
			want: 1,
		},
		{
			name: "Input: n = 2, pick = 2, Output: 2",
			args: args{
				n:    2,
				pick: 2,
			},
			want: 2,
		},
		{
			name: "Input: n = 10, pick = 6, Output: 6",
			args: args{
				n:    10,
				pick: 6,
			},
			want: 6,
		},
		{
			name: "Input: n = 2, pick = 1, Output: 1",
			args: args{
				n:    2,
				pick: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pick = tt.args.pick
			assert.Equalf(t, tt.want, guessNumber(tt.args.n), "guessNumber(%v)", tt.args.n)
		})
	}
}
