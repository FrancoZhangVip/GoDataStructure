package hot_one_hundred

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	incr := 0
	p, q, r := l1, l2, head
	for p != nil && q != nil {

		node := new(ListNode)
		node.Val = (p.Val + q.Val + incr) % 10

		if (p.Val + q.Val + incr) > 9 {
			incr = 1
		} else {
			incr = 0
		}

		r.Next = node
		p = p.Next
		q = q.Next
		r = r.Next

	}
	for p != nil {

		node := new(ListNode)
		node.Val = (p.Val + incr) % 10
		if (p.Val + incr) > 9 {
			incr = 1
		} else {
			incr = 0
		}

		r.Next = node
		r = r.Next
		p = p.Next

	}

	for q != nil {

		node := new(ListNode)
		node.Val = (q.Val + incr) % 10
		if (q.Val + incr) > 9 {
			incr = 1
		} else {
			incr = 0
		}

		r.Next = node
		r = r.Next
		q = q.Next

	}
	if incr == 1 {
		node := new(ListNode)
		node.Val = 1
		r.Next = node
	}
	return head.Next
}
