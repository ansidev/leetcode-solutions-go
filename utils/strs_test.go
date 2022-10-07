package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeProblemInput(t *testing.T) {
	tests := []struct {
		s    []string
		want string
	}{
		{
			s:    []string{"704"},
			want: "0704_binary_search",
		},
		{
			s:    []string{"35"},
			want: "0035_search_insert_position",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Input: %s, Output: %s", tt.s, tt.want), func(t *testing.T) {
			got := NormalizeProblemInput(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
