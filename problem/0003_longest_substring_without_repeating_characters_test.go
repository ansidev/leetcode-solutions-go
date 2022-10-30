package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type sargs struct {
	s string
}

func getLengthOfLongestSubstringTestData() []struct {
	name string
	args sargs
	want int
} {
	return []struct {
		name string
		args sargs
		want int
	}{
		{
			name: "Input: au, Output: 2",
			args: sargs{
				s: "au",
			},
			want: 2,
		},
		{
			name: "Input: bacabcbb, Output: 3",
			args: sargs{
				s: "bacabcbb",
			},
			want: 3,
		},
		{
			name: "Input: abcabcbb, Output: 3",
			args: sargs{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "Input: bbbbb, Output: 1",
			args: sargs{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "Input: pwwkew, Output: 3",
			args: sargs{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "Input: dvdf, Output: 3",
			args: sargs{
				s: "dvdf",
			},
			want: 3,
		},
	}
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := getLengthOfLongestSubstringTestData()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLongestSubstring(tt.args.s), "lengthOfLongestSubstring(%v)", tt.args.s)
		})
	}
}

func Test_lengthOfLongestSubstring2(t *testing.T) {
	tests := getLengthOfLongestSubstringTestData()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLongestSubstring2(tt.args.s), "lengthOfLongestSubstring(%v)", tt.args.s)
		})
	}
}
