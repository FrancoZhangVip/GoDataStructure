package list_structure

import "fmt"

type ListNode struct {
	val  int
	pre  *ListNode
	next *ListNode
}

type LinkList struct {
	head *ListNode
	tail *ListNode
	size int
}

func NewListNode(val int) (node *ListNode) {
	node = new(ListNode)
	node.val = val
	return node
}

func NewLinkList() (linkList *LinkList) {
	linkList = new(LinkList)
	return linkList
}

func (linkList *LinkList) AddNode(node *ListNode) {
	if linkList.size == 0 {
		linkList.head = node
		linkList.tail = node
	} else {
		linkList.tail.next = node
		node.pre = linkList.tail
		linkList.tail = node
	}

	linkList.size++
}

func (linkList *LinkList) DelNode(node *ListNode) {
	if linkList.size <= 0 {
		return
	}
	if linkList.size == 1 && linkList.head.val == node.val {
		linkList.head = nil
		linkList.tail = nil
		linkList.size--
		return
	}
	for i := linkList.head; i != nil; i = i.next {
		if i.val == node.val {
			if i.pre != nil {
				i.pre.next = i.next
			} else {
				linkList.head = i.next
			}

			if i.next != nil {
				i.next.pre = i.pre
			} else {
				linkList.tail = i.pre
			}

			linkList.size--
			return
		}
	}
}

func (linkList *LinkList) Size() int {
	return linkList.size
}

func (linkList *LinkList) Show() {
	for i := linkList.head; i != nil; i = i.next {
		fmt.Print(i.val, "\t")
	}
	fmt.Println()
}
