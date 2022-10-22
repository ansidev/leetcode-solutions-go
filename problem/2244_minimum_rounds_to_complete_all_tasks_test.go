package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumRounds(t *testing.T) {
	type args struct {
		tasks []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Input: tasks = [2,2,3,3,2,4,4,4,4,4], Output: 4",
			args: args{
				tasks: []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4},
			},
			want: 4,
		},
		{
			name: "Input: tasks = [2,3,3], Output: -1",
			args: args{
				tasks: []int{2, 3, 3},
			},
			want: -1,
		},
		{
			name: "Input: tasks = [7,7,7,7,7,7], Output: 2",
			args: args{
				tasks: []int{7, 7, 7, 7, 7, 7},
			},
			want: 2,
		},
		{
			name: "Input: tasks = [7,7,7,7,7,7,7,7,7,7,7,7,7], Output: 5",
			args: args{
				tasks: []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumRounds(tt.args.tasks), "minimumRounds(%v)", tt.args.tasks)
		})
	}
}
