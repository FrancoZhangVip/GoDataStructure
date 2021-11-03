package list

/*
题意：

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

示例：

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3

链接：https://leetcode-cn.com/problems/design-linked-list


*/

type MyLinkedList struct {
	Head *ListNode
	Size int
}

func Constructor() MyLinkedList {
	var list MyLinkedList
	return list
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Size || index < 0 {
		return -1
	}

	pointer := this.Head
	for i := 0; i < index; i++ {
		pointer = pointer.Next
	}
	return pointer.Val

}

func (this *MyLinkedList) AddAtHead(val int) {
	head := new(ListNode)
	head.Val = val
	head.Next = this.Head
	this.Head = head
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	tail := new(ListNode)
	tail.Val = val
	if this.Head == nil {
		this.Head = tail
		this.Size++
		return
	}
	pointer := this.Head
	for pointer.Next != nil {
		pointer = pointer.Next
	}
	pointer.Next = tail
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.Size {
		this.AddAtTail(val)
		return
	}
	if index > this.Size {
		return
	}
	node := new(ListNode)
	node.Val = val

	pointer := this.Head
	for i := 0; i < index-1; i++ {
		pointer = pointer.Next
	}
	node.Next = pointer.Next
	pointer.Next = node
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.Size-1 < index || index < 0 {
		return
	}
	hiden := new(ListNode)
	hiden.Next = this.Head

	pointer := hiden
	for i := 0; i < index; i++ {
		pointer = pointer.Next
	}
	if pointer.Next != nil {
		pointer.Next = pointer.Next.Next
	} else {
		pointer.Next = nil
	}

	this.Head = hiden.Next
	this.Size--
}

func (this *MyLinkedList) Show() {
	head := this.Head
	for head != nil {
		print(head.Val, "->")
		head = head.Next
	}
	println()

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
