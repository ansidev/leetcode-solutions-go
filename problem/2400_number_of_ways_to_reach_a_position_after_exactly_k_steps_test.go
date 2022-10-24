package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type numberOfWaysArgs struct {
	startPos int
	endPos   int
	k        int
}

var numberOfWaysTests = []struct {
	name string
	args numberOfWaysArgs
	want int
}{
	{
		name: "Input: startPos = 1, endPos = 2, k = 3, Output: 3",
		args: numberOfWaysArgs{
			startPos: 1,
			endPos:   2,
			k:        3,
		},
		want: 3,
	},
	{
		name: "Input: startPos = 2, endPos = 5, k = 10, Output: 0",
		args: numberOfWaysArgs{
			startPos: 2,
			endPos:   5,
			k:        10,
		},
		want: 0,
	},
	{
		name: "Input: startPos = 1, endPos = 1000, k = 999, Output: 1",
		args: numberOfWaysArgs{
			startPos: 1,
			endPos:   1000,
			k:        999,
		},
		want: 1,
	},
	{
		name: "Input: startPos = 989, endPos = 1000, k = 99, Output: 934081896",
		args: numberOfWaysArgs{
			startPos: 989,
			endPos:   1000,
			k:        99,
		},
		want: 934081896,
	},
	{
		name: "Input: startPos = 999, endPos = 1000, k = 1000, Output: 0",
		args: numberOfWaysArgs{
			startPos: 999,
			endPos:   1000,
			k:        1000,
		},
		want: 0,
	},
	{
		name: "Input: startPos = 1, endPos = 3, k = 1, Output: 0",
		args: numberOfWaysArgs{
			startPos: 1,
			endPos:   3,
			k:        1,
		},
		want: 0,
	},
}

func Test_numberOfWays1(t *testing.T) {
	for _, tt := range numberOfWaysTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numberOfWays1(tt.args.startPos, tt.args.endPos, tt.args.k), "numberOfWays1(%v, %v, %v)", tt.args.startPos, tt.args.endPos, tt.args.k)
		})
	}
}

func Test_numberOfWays2(t *testing.T) {
	for _, tt := range numberOfWaysTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numberOfWays2(tt.args.startPos, tt.args.endPos, tt.args.k), "numberOfWays2(%v, %v, %v)", tt.args.startPos, tt.args.endPos, tt.args.k)
		})
	}
}
