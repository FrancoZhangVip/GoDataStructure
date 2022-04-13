package leetcode_algorithm

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	hide := new(ListNode)
	hide.Next = head
	p, q := hide, hide
	for i := 0; i < n+1; i++ {
		q = q.Next
	}
	for true {
		if q == nil {
			break
		} else {
			p = p.Next
			q = q.Next
		}
	}
	p.Next = p.Next.Next
	return hide.Next
}
