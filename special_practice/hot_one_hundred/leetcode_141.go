package hot_one_hundred

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	p, q := head, head
	for true {
		q = q.Next
		if q == p {
			return true
		}
		if q == nil {
			return false
		}
		q = q.Next
		if q == p {
			return true
		}
		if q == nil {
			return false
		}
		p = p.Next
		if q == p {
			return true
		}
		if p == nil {
			return false
		}
	}
	return false
}
