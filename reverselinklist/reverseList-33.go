package reverselinklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next = prev
		prev = head
		head = head.Next
	}
	return prev
}
