package problem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Input: l1 = [2], l2 = [5], Output: [7]",
			args: args{
				l1: &ListNode{
					Val:  2,
					Next: nil,
				},
				l2: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  7,
				Next: nil,
			},
		},
		{
			name: "Input: l1 = [2], l2 = [8], Output: [0, 1]",
			args: args{
				l1: &ListNode{
					Val:  2,
					Next: nil,
				},
				l2: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
			want: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
		{
			name: "Input: l1 = [2,6], l2 = [5,3], Output: [7, 9]",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  6,
						Next: nil,
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val:  9,
					Next: nil,
				},
			},
		},
		{
			name: "Input: l1 = [2,4,3], l2 = [5,6,4], Output: 807",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "Input: l1 = [9,9,9], l2 = [9,9,9,9], Output: 10998",
			args: args{
				l1: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  9,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, addTwoNumbers(tt.args.l1, tt.args.l2), "addTwoNumbers(%v, %v)", tt.args.l1, tt.args.l2)
		})
	}
}
