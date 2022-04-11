package sword_towards_offer

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	copyMap := make(map[*Node]*Node)

	var newHead *Node
	var pre *Node

	first := true
	for p := head; p != nil; p = p.Next {
		q := new(Node)
		q.Val = p.Val
		copyMap[p] = q

		if first {
			newHead = q
			pre = q
			first = false
		} else {
			pre.Next = q
			pre = q
		}

	}

	p := head
	q := newHead
	for true {
		if p == nil {
			break
		} else if p.Random == nil {
			q.Random = nil
		} else {
			q.Random = copyMap[p.Random]
		}
		p = p.Next
		q = q.Next
	}
	return newHead
}
