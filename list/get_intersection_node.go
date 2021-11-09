package list

/*
给定两个（单向）链表，判定它们是否相交并返回交点。请注意相交的定义基于节点的引用，而不是基于节点的值。换句话说，如果一个链表的第k个节点与另一个链表的第j个节点是同一节点（引用完全相同），则这两个链表相交。

示例 1：

输入：listA = [4,1,8,4,5], listB = [5,0,1,8,4,5]

输出：Reference of the node with value = 8

输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	counterA := 0
	counterB := 0
	qa := headA
	qb := headB
	for true {

		if qa == nil && counterA == 0 {
			counterA++
			qa = headB
		}
		if qb == nil && counterB == 0 {
			counterB++
			qb = headA
		}
		if qa == nil && counterA == 1 {
			return nil
		}
		if qb == nil && counterB == 1 {
			return nil
		}
		if qa == qb {
			return qa
		}
		qa = qa.Next
		qb = qb.Next
	}
	return nil
}
