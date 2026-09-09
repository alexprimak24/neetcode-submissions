/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		nextCurr := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextCurr
	}

	return prev

}
