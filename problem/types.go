package problem

func NewNode(val int) *Node {
	return &Node{Val: val}
}

type Node struct {
	Val      int
	Children []*Node
}

func NewListNode() *ListNode {
	return &ListNode{}
}

func NewListNodeWithValue(val int) *ListNode {
	return &ListNode{Val: val}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Prepend(num int) *ListNode {
	l.Next = l
	l.Val = num
	return l
}

func (l *ListNode) Append(num int) *ListNode {
	if l == nil {
		return NewListNodeWithValue(num)
	}

	tmp := l
	for tmp.Next != nil {
		tmp = tmp.Next
	}

	tmp.Next = NewListNodeWithValue(num)

	return l
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
