package sword_towards_offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	q := head.Next
	head.Next = nil

	for true {
		r := q.Next
		q.Next = p
		p = q
		q = r
		if q == nil {
			break
		}
	}
	return p
}
