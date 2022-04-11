package sword_towards_offer

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

func ReversePrint(head *ListNode) []int {

	counter := 0
	for index := head; index != nil; index = index.Next {
		counter++
	}
	res := make([]int, counter)
	cur := head
	for i := counter - 1; i >= 0; i-- {
		res[i] = cur.Val
		cur = cur.Next
	}
	return res
}
