package main

func main() {}

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// 203. 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	virHead := &ListNode{Next: head}
	curNode := virHead
	for curNode.Next != nil {
		if curNode.Next.Val == val {
			curNode.Next = curNode.Next.Next
		} else {
			curNode = curNode.Next
		}
	}
	return virHead.Next
}

type MyLinkedList struct {
	virHead *ListNode
	size    int
}

func Constructor() MyLinkedList {
	return MyLinkedList{virHead: &ListNode{}, size: 0}
}

func (this *MyLinkedList) Get(index int) int {
	node := this.getNode(index)
	if node == nil {
		return -1
	}
	return node.Val
}

func (this *MyLinkedList) getNode(index int) *ListNode {
	if index < 0 || index >= this.size {
		return nil
	}
	curNode := this.virHead
	for i := 0; i <= index; i++ {
		curNode = curNode.Next
	}
	return curNode
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.virHead.Next = &ListNode{Val: val, Next: this.virHead.Next}
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{Val: val}
	node := this.getNode(this.size - 1)
	defer func() {
		this.size++
	}()
	if node == nil {
		this.virHead.Next = newNode
		return
	}
	node.Next = newNode
}

func (this *MyLinkedList) AddAtIndex(index, val int) {
	if index > this.size {
		return
	}
	cur := this.virHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	if index == this.size {
		cur.Next = &ListNode{Val: val}
	} else {
		cur.Next = &ListNode{Val: val, Next: cur.Next}
	}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	cur := this.virHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	this.size--
}

// 206. 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre, cur = cur, tmp
	}
	return pre
}
