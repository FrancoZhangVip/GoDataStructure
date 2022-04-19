package leetcode_algorithm

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := new(ListNode)
	cur := res
	p := list1
	q := list2
	for true {
		if p == nil || q == nil {
			break
		} else if p.Val < q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}

	for p != nil {
		cur.Next = p
		cur = cur.Next
		p = p.Next
	}

	for q != nil {
		cur.Next = q
		cur = cur.Next
		q = q.Next
	}
	return res.Next
}
