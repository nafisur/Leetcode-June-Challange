package leetcode

/*
Delete Node in a Linked List
Solution
Write a function to delete a node (except the tail) in a singly linked list, given only access to that node.

Given linked list -- head = [4,5,1,9], which looks like following:

Example 1:

Input: head = [4,5,1,9], node = 5
Output: [4,1,9]

Example 2:

Input: head = [4,5,1,9], node = 1
Output: [4,5,9]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	if node.Next == nil {
		return
	}

	node.Val = node.Next.Val
	node.Next = node.Next.Next

}
