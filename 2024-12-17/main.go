package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 24. 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	virHead := &ListNode{Next: head}
	curNode := virHead
	// 两个节点都存在，则互换
	for curNode.Next != nil && curNode.Next.Next != nil {
		tmp := curNode.Next
		tmp2 := curNode.Next.Next.Next
		curNode.Next = curNode.Next.Next
		curNode.Next.Next = tmp
		curNode.Next.Next.Next = tmp2
		curNode = curNode.Next.Next

	}
	return virHead.Next
}

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	virHead := &ListNode{Next: head}
	slow := virHead
	fast := virHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return virHead.Next
}
