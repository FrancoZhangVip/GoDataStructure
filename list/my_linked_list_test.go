package list

import "testing"

func TestMyLinkedList(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(2)
	obj.DeleteAtIndex(1)
	obj.AddAtHead(2)
	obj.AddAtHead(7)
	obj.AddAtHead(3)
	obj.AddAtHead(2)
	obj.AddAtHead(5)
	obj.AddAtTail(5)
	println(obj.Get(5))
	obj.DeleteAtIndex(6)
	obj.DeleteAtIndex(4)
}

func TestMyLinkedList2(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2) //链表变为1-> 2-> 3
	println(obj.Get(1))  //返回2
	obj.DeleteAtIndex(1) //现在链表是1-> 3
	println(obj.Get(1))  //返回3
}

func TestMyLinkedList3(t *testing.T) {
	function := []string{"MyLinkedList", "addAtHead", "addAtIndex", "addAtIndex", "addAtHead", "deleteAtIndex", "addAtIndex", "addAtHead", "addAtTail", "addAtHead", "get", "addAtTail", "addAtTail", "addAtIndex", "addAtTail", "addAtTail", "addAtHead", "addAtTail", "addAtHead", "addAtTail", "addAtHead", "addAtTail", "addAtTail", "addAtHead", "addAtTail", "addAtIndex", "addAtHead", "deleteAtIndex", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtIndex", "addAtHead", "addAtTail", "addAtHead", "deleteAtIndex", "addAtTail", "deleteAtIndex", "addAtTail", "addAtTail", "addAtTail", "addAtTail", "addAtHead", "addAtTail", "get", "addAtIndex", "get", "deleteAtIndex", "addAtTail", "addAtHead", "addAtTail", "addAtTail", "addAtIndex", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtTail", "deleteAtIndex", "deleteAtIndex", "addAtHead", "addAtHead", "addAtTail", "addAtHead", "get", "addAtIndex", "addAtIndex", "get", "addAtTail", "get", "addAtTail", "deleteAtIndex", "get", "addAtTail", "addAtTail", "addAtHead", "addAtTail", "deleteAtIndex", "addAtHead", "addAtHead", "addAtHead", "deleteAtIndex", "addAtTail", "addAtIndex", "addAtTail", "addAtTail", "addAtIndex", "addAtIndex", "addAtHead", "addAtIndex", "addAtHead", "addAtHead", "addAtTail", "addAtIndex", "addAtTail", "get", "addAtHead", "addAtTail", "addAtHead", "addAtHead"}
	value := [][]int{{}, {86}, {1, 54}, {1, 14}, {83}, {4}, {3, 18}, {46}, {3}, {76}, {5}, {79}, {74}, {7, 28}, {68}, {16}, {82}, {24}, {30}, {96}, {21}, {79}, {66}, {2}, {2}, {7, 57}, {59}, {21}, {19}, {55}, {46}, {17}, {16, 41}, {97}, {85}, {63}, {30}, {11}, {16}, {91}, {29}, {47}, {20}, {12}, {80}, {15}, {12, 8}, {21}, {30}, {11}, {54}, {51}, {41}, {8, 88}, {62}, {7}, {59}, {73}, {73}, {55}, {9}, {7}, {49}, {99}, {17}, {44}, {11}, {26, 86}, {10, 99}, {19}, {71}, {29}, {32}, {2}, {3}, {16}, {2}, {83}, {31}, {3}, {23}, {64}, {96}, {27}, {81}, {12, 78}, {21}, {69}, {5, 35}, {8, 72}, {60}, {19, 73}, {55}, {83}, {86}, {31, 70}, {49}, {19}, {64}, {22}, {25}, {13}}

	var obj MyLinkedList
	for i := 0; i < len(function); i++ {
		switch function[i] {
		case "MyLinkedList":
			obj = Constructor()
		case "addAtHead":
			obj.AddAtHead(value[i][0])
		case "addAtIndex":
			obj.AddAtIndex(value[i][0], value[i][1])
		case "AddAtTail":
			obj.AddAtTail(value[i][0])
		case "get":
			println(obj.Get(value[i][0]))
		case "deleteAtIndex":
			obj.DeleteAtIndex(value[i][0])
		}
	}

}

func TestMyLinkedList4(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtHead(3)
	obj.AddAtHead(2)
	obj.AddAtHead(4)
	obj.AddAtHead(5) //5 4 2 3 1
	obj.Show()
}

func TestMyLinkedList5(t *testing.T) {
	obj := Constructor()
	obj.AddAtTail(1)
	obj.AddAtTail(3)
	obj.AddAtTail(2)
	obj.AddAtTail(4)
	obj.AddAtTail(5) //1 3 2 4 5
	obj.Show()
}

func TestMyLinkedList6(t *testing.T) {
	obj := Constructor()
	obj.AddAtTail(1)
	obj.AddAtTail(3)
	obj.AddAtTail(2)
	obj.AddAtTail(4)
	obj.AddAtTail(5) //1 3 2 4 5

	obj.AddAtIndex(-1, 7)
	obj.AddAtIndex(5, 8)
	obj.AddAtIndex(3, 9)
	obj.AddAtIndex(3, 10) //7 1 3 10 9 2 4 8 5
	obj.Show()
}

func TestMyLinkedList7(t *testing.T) {
	obj := Constructor()
	obj.AddAtTail(1)
	obj.AddAtTail(3)
	obj.AddAtTail(2)
	obj.AddAtTail(4)
	obj.AddAtTail(5) //1 3 2 4 5

	obj.DeleteAtIndex(0)
	obj.DeleteAtIndex(-1)
	obj.DeleteAtIndex(5) //3 2 4 5
	obj.Show()
	obj.DeleteAtIndex(3) //3 2 4
	obj.Show()
	obj.DeleteAtIndex(3) //3 2 4
	obj.Show()
}
