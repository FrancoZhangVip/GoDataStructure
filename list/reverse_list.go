package list

/*
题意：反转一个单链表。

示例: 输入: 1->2->3->4->5->NULL 输出: 5->4->3->2->1->NULL


https://leetcode-cn.com/problems/reverse-linked-list/
*/

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	hiden := new(ListNode)
	hiden.Next = head

	quick := head
	slow := hiden
	for true {
		temp := quick.Next
		quick.Next = slow
		slow = quick
		quick = temp
		if quick == nil {
			break
		}
	}
	head.Next = nil
	return slow
}

/*
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}


链接：https://leetcode-cn.com/problems/reverse-linked-list/solution/fan-zhuan-lian-biao-by-leetcode-solution-d1k2/

*/
