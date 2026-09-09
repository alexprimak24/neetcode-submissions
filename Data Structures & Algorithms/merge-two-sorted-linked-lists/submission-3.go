/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	merged := dummy

	pt1 := list1
	pt2 := list2

	for pt1 != nil && pt2 != nil {
		if pt1.Val < pt2.Val {
			merged.Next = pt1
			pt1 = pt1.Next
		} else {
			merged.Next = pt2
			pt2 = pt2.Next
		}
		merged = merged.Next
	}

	merged.Next = pt1
	if pt1 == nil {
		merged.Next = pt2
	}

	return dummy.Next
}
