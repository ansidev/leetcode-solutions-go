package problem

import (
	"strconv"
	"strings"
)

func splitMessage(message string, limit int) []string {
	mLen := len(message)

	b := 1
	aLen := sz(1)

	for b*limit < b*(sz(b)+3)+aLen+mLen {
		if sz(b)*2+3 >= limit {
			return []string{}
		}
		b++
		aLen = aLen + sz(b)
	}

	rs := make([]string, 0)
	i := 0
	for a := 1; a <= b; a++ {
		var sb strings.Builder
		j := limit - (3 + sz(a) + sz(b))
		sb.WriteString(message[i:min(i+j, mLen)])
		sb.WriteRune('<')
		sb.WriteString(strconv.Itoa(a))
		sb.WriteRune('/')
		sb.WriteString(strconv.Itoa(b))
		sb.WriteRune('>')
		rs = append(rs, sb.String())
		i = i + j
	}

	return rs
}

func sz(i int) int {
	return len(strconv.Itoa(i))
}
